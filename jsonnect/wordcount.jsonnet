local workflow = import "workflow.libsonnet";

// Workflow that performs a wordcount using shell commands.
{
  wordcount: workflow.Workflow {
    retries: 12,
    schedule: workflow.Schedule {
      start_date: "2015-11-15",
      start_time: "17:30",
      repeat_frequency: 1,
      repeat_type: "week",
    },
    jobs: {
      local input_file = "/tmp/passage_test",
      local tokens_file = "/tmp/tokens",
      local sorted_tokens_file = "/tmp/sorted_tokens",
      local counts_file = "/tmp/counts",

      // Reads the input file and produces an output file with one word per
      // line.
      tokenize: workflow.ShJob {
        command: "tr ' ' '\n' < %s > %s" % [input_file, tokens_file],
        inputs: [input_file],
        outputs: [tokens_file],
      },

      // Takes the tokens file and produces a file with the tokens sorted.
      sort: workflow.ShJob {
        deps: [":tokenize"],
        command: "sort %s > %s" % [tokens_file, sorted_tokens_file],
        inputs: [tokens_file],
        outputs: [sorted_tokens_file],
      },

      // Takes the file containing sorted tokens and produces a file containing
      // the counts for each word.
      count: workflow.ShJob {
        deps: [":sort"],
        command: "uniq -c %s > %s" % [sorted_tokens_file, counts_file],
        inputs: [sorted_tokens_file],
        outputs: [counts_file],
      },
    }
  }
}
