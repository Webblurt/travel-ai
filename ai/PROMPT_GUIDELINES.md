# AI Development Workflow

This repository is developed with an **AI-first approach**.

### Core principles:
- The developer **commands** AI, AI does not "suggest".
- Every feature starts from a prompt stored in `/ai/prompts`.
- AI-generated code is **reviewed manually** for architectural consistency.
- If AI output is incorrect, feedback is added back to the prompt template.

### Prompt types:
- `/ai/prompts/feature_request.md` - how features are described.
- `/ai/prompts/codegen_frontend.md` - how to request SvelteKit code.
- `/ai/prompts/codegen_backend.md` - how to request Go code.

Logs of AI interactions are stored in `/ai/logs` to simulate decision trace.
