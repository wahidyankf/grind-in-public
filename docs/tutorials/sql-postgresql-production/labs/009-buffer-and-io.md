---
tldr: "Simulates clock-sweep eviction and separates cache, scan, and storage symptoms."
when_to_use: "Use after buffer-cache and I/O internals."
---

# Lab 009: Buffer Replacement and I/O

Four buffers have usage counts `[2, 0, 1, 0]`, with the clock hand at index 0. Choose the next eviction candidate,
showing decrements. Explain why a large sequential scan uses a ring and what buffer/I/O metrics distinguish cache miss
from lock wait.

Compare with the [solution](../lab-solutions/009-buffer-and-io.md).
