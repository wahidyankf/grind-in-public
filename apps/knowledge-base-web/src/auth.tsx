import {
  createContext,
  useContext,
  useMemo,
  useState,
  type ReactNode,
} from "react";
import type { PublicUser } from "./api";

export type Session = { accessToken: string; user: PublicUser };

type AuthValue = {
  session: Session | null;
  signIn: (session: Session) => void;
};

const AuthContext = createContext<AuthValue | null>(null);

// The token lives in memory only. A page reload signs the user out, which is accepted for this first step.
export function AuthProvider({ children }: { children: ReactNode }) {
  const [session, setSession] = useState<Session | null>(null);
  const value = useMemo(() => ({ session, signIn: setSession }), [session]);
  return <AuthContext value={value}>{children}</AuthContext>;
}

export function useAuth(): AuthValue {
  const value = useContext(AuthContext);
  if (!value) {
    throw new Error("useAuth must be used inside AuthProvider");
  }
  return value;
}
