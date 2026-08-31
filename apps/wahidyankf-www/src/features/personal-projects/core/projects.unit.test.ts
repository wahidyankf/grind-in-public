import { describe, it, expect } from "vitest";
import { projects, filterProjects, type Project } from "./projects";

// `filterProjects` is a thin wrapper over the generic `filterItems`, and its
// value is entirely in which fields it names as searchable. These tests assert
// that list, field by field, so removing one from the wrapper fails here rather
// than quietly narrowing what the projects page can find.
const sample: Project = {
  title: "Sample Project",
  description: "A description mentioning cartography",
  period: "January 2020 - Present",
  details: ["A detail mentioning telemetry"],
  skills: ["Facilitation"],
  programmingLanguages: ["Elixir"],
  frameworks: ["Phoenix"],
  aiSkills: ["Prompting"],
  links: { repository: "https://example.test/repo" },
};

describe("projects record", () => {
  it("ships at least one project", () => {
    expect(projects.length).toBeGreaterThan(0);
  });

  it("gives every project a title and a period", () => {
    for (const project of projects) {
      expect(project.title.length).toBeGreaterThan(0);
      expect(project.period.length).toBeGreaterThan(0);
    }
  });
});

describe("filterProjects", () => {
  it("returns every project when the term is empty", () => {
    expect(filterProjects(projects, "")).toEqual(projects);
  });

  it("returns nothing when no project matches", () => {
    expect(filterProjects([sample], "nonexistent-term")).toEqual([]);
  });

  it("matches case-insensitively", () => {
    expect(filterProjects([sample], "SAMPLE")).toEqual([sample]);
  });

  it.each([
    ["title", "Sample Project"],
    ["description", "cartography"],
    ["details", "telemetry"],
    ["skills", "Facilitation"],
    ["programmingLanguages", "Elixir"],
    ["frameworks", "Phoenix"],
    ["aiSkills", "Prompting"],
  ])("searches the %s field", (_field, term) => {
    expect(filterProjects([sample], term)).toEqual([sample]);
  });

  it("does not search fields outside that list", () => {
    // `period` and `links` are deliberately not searchable: a reader looking
    // for "Phoenix" wants the framework, not every project whose repository
    // URL happens to contain the word.
    expect(filterProjects([sample], "example.test")).toEqual([]);
    expect(filterProjects([sample], "January 2020")).toEqual([]);
  });
});
