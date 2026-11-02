export type PublicUser = { id: string; email: string };

export type LoginResponse = {
  accessToken: string;
  tokenType: "bearer";
  expiresIn: number;
  user: PublicUser;
};

export class ApiError extends Error {
  readonly status: number;

  constructor(status: number, message: string) {
    super(message);
    this.name = "ApiError";
    this.status = status;
  }
}

async function postJson<T>(path: string, body: unknown): Promise<T> {
  const response = await fetch(path, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(body),
  });
  const data: unknown = await response.json().catch(() => null);
  if (!response.ok) {
    const message =
      typeof data === "object" &&
      data !== null &&
      "message" in data &&
      typeof data.message === "string"
        ? data.message
        : "Request failed";
    throw new ApiError(response.status, message);
  }
  return data as T;
}

export function registerUser(
  email: string,
  password: string,
): Promise<PublicUser> {
  return postJson<PublicUser>("/api/auth/register", { email, password });
}

export function loginUser(
  email: string,
  password: string,
): Promise<LoginResponse> {
  return postJson<LoginResponse>("/api/auth/login", { email, password });
}
