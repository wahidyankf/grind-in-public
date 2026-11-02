import { SignJWT, jwtVerify } from "jose";

const ALGORITHM = "HS256";

export function signAccessToken(
  userId: string,
  secret: string,
  ttlSeconds: number,
): Promise<string> {
  return new SignJWT({})
    .setProtectedHeader({ alg: ALGORITHM })
    .setSubject(userId)
    .setIssuedAt()
    .setExpirationTime(`${ttlSeconds}s`)
    .sign(new TextEncoder().encode(secret));
}

// Returns the user id from a valid token. Throws if the signature or expiry is wrong.
export async function verifyAccessToken(
  token: string,
  secret: string,
): Promise<string> {
  const { payload } = await jwtVerify(token, new TextEncoder().encode(secret), {
    algorithms: [ALGORITHM],
  });
  if (!payload.sub) {
    throw new Error("Token has no subject");
  }
  return payload.sub;
}
