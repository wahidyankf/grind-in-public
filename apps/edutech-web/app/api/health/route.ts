// Playwright and load balancers call this to know the server is ready.
export function GET() {
  return Response.json({ status: "ok" });
}
