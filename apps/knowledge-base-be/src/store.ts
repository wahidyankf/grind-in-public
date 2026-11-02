export type PublicUser = { id: string; email: string };
export type StoredUser = PublicUser & { passwordHash: string };
export type NewUser = { email: string; passwordHash: string };

export interface AuthStore {
  // Returns null when the email is already registered.
  createUser(user: NewUser): Promise<PublicUser | null>;
  findUserByEmail(email: string): Promise<StoredUser | null>;
}
