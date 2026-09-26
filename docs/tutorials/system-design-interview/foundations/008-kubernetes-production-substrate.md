---
tldr:
  "Maps application requirements to Kubernetes workloads, probes, resources, scaling, disruption, and rollout controls."
when_to_use: "Use when deploying Python services on Kubernetes or discussing migration and reliability trade-offs."
---

# Kubernetes as a Production Substrate

Kubernetes schedules and reconciles workloads; it does not create application idempotency, safe database migrations, or
useful service boundaries. Start with the application's state and failure semantics, then map them to primitives.

## Runtime mapping

```text
external client
      |
  [Ingress/Gateway]
      |
   [Service] ---------------- stable discovery
      |
  +---+-------------------+
  | Deployment            |
  | pod A  pod B  pod C   |--- ConfigMap / Secret references
  +---+-------------------+
      |
 managed database / durable log / object storage
```

- `Deployment`: interchangeable stateless replicas and rolling releases.
- `StatefulSet`: stable identity/storage and ordered management when the application needs them.
- `Job`: bounded work that must complete; `CronJob`: scheduled job creation.
- `DaemonSet`: one workload per selected node, commonly infrastructure agents.
- `Service`: stable virtual endpoint for changing Pods.

Prefer managed durable state outside the cluster unless the team can operate quorum, backup, restore, upgrades, and
failure repair for that database. A persistent volume alone is not a database availability strategy.

## Complete deployment slice

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: decision-api
spec:
  replicas: 3
  strategy:
    rollingUpdate:
      maxUnavailable: 0
      maxSurge: 1
  selector:
    matchLabels:
      app: decision-api
  template:
    metadata:
      labels:
        app: decision-api
    spec:
      terminationGracePeriodSeconds: 30
      containers:
        - name: api
          image: registry.example/decision-api:2026-09-26.1
          ports:
            - name: http
              containerPort: 8080
          resources:
            requests:
              cpu: "500m"
              memory: "512Mi"
            limits:
              memory: "1Gi"
          startupProbe:
            httpGet:
              path: /health/startup
              port: http
            failureThreshold: 30
            periodSeconds: 2
          readinessProbe:
            httpGet:
              path: /health/ready
              port: http
            periodSeconds: 5
          livenessProbe:
            httpGet:
              path: /health/live
              port: http
            periodSeconds: 10
          lifecycle:
            preStop:
              exec:
                command: ["/bin/sh", "-c", "sleep 5"]
```

The startup probe protects slow initialization from liveness restarts. Readiness removes a Pod that cannot serve new
requests. Liveness should detect an unrecoverable stuck process, not dependency health; restarting every Pod because a
shared database is down amplifies the incident.

Use immutable image tags or digests. Set CPU requests for scheduling and autoscaling. A CPU limit can throttle latency-
sensitive Python services, so add it only from measured policy; a memory limit provides an explicit OOM boundary.

## Graceful termination

```text
Pod termination requested
   |
   +-> readiness false -> endpoints converge
   +-> preStop / SIGTERM -> stop accepting -> drain in-flight work
   +-> flush bounded telemetry/checkpoints
   `-> exit before grace period; otherwise SIGKILL
```

Workers must stop fetching new messages, finish or safely abandon the current unit, and avoid acknowledging before the
effect is durable.

## Scaling

Horizontal Pod Autoscaler can scale on CPU, memory, or custom metrics. For APIs, concurrency or latency may be better
signals; for workers, oldest-message age and backlog per effective worker are often better than CPU.

```text
backlog rises -> HPA asks for Pods -> scheduler needs nodes -> node autoscaler adds capacity -> Pods warm
```

This loop may take minutes. Preserve headroom and include startup/model-load time. Never let replica scaling create more
database connections than the database can serve.

## Availability controls

- topology spread or anti-affinity distributes replicas across zones/nodes;
- Pod disruption budgets limit voluntary simultaneous disruption, not involuntary failure;
- priority and quotas protect critical workloads and tenants;
- network policy restricts reachable peers but requires tested DNS and egress rules;
- service accounts and workload identity minimize credential scope.

## Rollout and migration

Canary by traffic slice, tenant cohort, or shadow execution. Compare latency, errors, decisions, and domain drift. A
Kubernetes rollback changes Pods, not an incompatible database mutation. Use expand/contract schemas and separate
state-migration control.

## References

- [Deployments](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
- [Horizontal Pod Autoscaling](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)
- [Pod lifecycle](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/)
- [Pod disruptions](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/)
- [Topology spread constraints](https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/)
