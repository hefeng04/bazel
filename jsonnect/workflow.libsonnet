{
  // Configuration for a workflow.
  Workflow:: {
    schedule: {},
    retries: 5,
    jobs: {},
  },

  // Scheduling configuration for a workflow.
  Schedule:: {
    start_date: "",
    start_time: "",
    repeat_frequency: 0,
    repeat_type: "",
  },

  // Base configuration for a Job in a workflow.
  Job:: {
    type: "base",
    deps: [],
    inputs: [],
    outputs: [],
  },

  // Configuration for a job that runs a shell command.
  ShJob:: self.Job {
    type: "sh",
    command: "",
    vars: {},
  }
}
