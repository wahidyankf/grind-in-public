export type BehaviourTestLayer = "unit" | "integration";

const configuredLayer = process.env.WAHIDYANKF_WWW_BEHAVIOUR_TEST_LAYER;

if (configuredLayer !== "unit" && configuredLayer !== "integration") {
  throw new Error(
    "WAHIDYANKF_WWW_BEHAVIOUR_TEST_LAYER must be 'unit' or 'integration'.",
  );
}

export const behaviourTestLayer: BehaviourTestLayer = configuredLayer;
