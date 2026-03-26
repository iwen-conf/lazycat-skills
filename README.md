# 🐱 Lazycat Skills (AI Skill Pack for LazyCat MicroServer)

This is a comprehensive set of **AI Agent Skills** specifically designed for developers on the **LazyCat MicroServer** platform. 
By integrating these skills, your AI assistants (such as Cursor, Windsurf, Cline, etc.) transform into Lazycat ecosystem experts, capable of automatically writing `lzc-manifest.yml`, packaging LPKs, handling routing, and implementing MicroServer authentication.

## 📦 Available Skills

The skill pack currently includes the following core and specialized skills:

### Core Lifecycle Skills
- `lazycat:ship-app`: **Overall Control**. Manages the entire lifecycle from project initiation, asset organization, packaging, submission for review, to post-release verification. Also handles AI App / AI Browser Extension delivery assessment.
- `lazycat:create-app`: **Project Initialization**. Responsible for new project creation, initial documentation structure, tech stack baselines, native capability integration, unified authentication, and AI configuration baselines.
- `lazycat:update-app`: **Application Updates**. Manages the process of updating existing Lazycat applications, ensuring version consistency and compatibility.
- `lazycat:port-app`: **Application Porting**. Focuses on porting existing applications to Lazycat, including selection, GitHub research, App Store de-duplication, script entry points, and AI Pod route assessment.

### Specialized Technical Skills
- `lazycat:developer-expert`: **All-around Developer Expert**. General-purpose expert for all things LazyCat MicroServer development.
- `lazycat:lpk-builder`: **Packaging Expert**. Specializes in transforming Docker/source code into `.lpk` Lazycat MicroServer applications.
- `lazycat:advanced-routing`: **Network & Routing**. Handles complex network requirements like multi-domain configuration, Layer 4 forwarding (ingress), URL prefix stripping, and custom Nginx proxies.
- `lazycat:auth-integration`: **Authentication & Identity**. Expert in Lazycat API access, OIDC login integration, and identity management.
- `lazycat:aipod-developer`: **AI Pod Development**. Specialized guide for developing applications utilizing LazyCat's AI Pod and compute power infrastructure.
- `lazycat:dynamic-deploy`: **Dynamic Deployment**. Handles the intricacies of Lazycat's dynamic deployment mechanisms.

### Assets & Documentation Skills
- `lazycat:admin-ui`: **Admin UI Quality**. Ensures admin interfaces, operational consoles, and B-side workspaces are high-quality, branded, and screenshot-ready.
- `lazycat:prepare-icon`: **Icon Preparation**. Generates high-quality App Icon prompts suitable for external image models during the asset preparation phase.
- `lazycat:write-guide`: **Guide & Documentation**. Focuses on writing high-quality application guides and tutorials following official incentive rules.

## 📂 Repository Structure

This repository follows the **Progressive Disclosure** principle for Agent loading:
```text
lazycat-skills/
├── README.md              # Project overview (Human-facing)
├── AGENTS.md              # AI Agent behavior constraints
├── lazycat:create-app/
│   ├── SKILL.md           # Core skill instructions and triggers
│   └── references/        # Detailed reference documentation
└── ... (other skills)
```

## 🚀 How to Use

We recommend using the `npx skills` tool to add these to your AI assistant's workspace:

```bash
# Execute in your project root:
npx skills add whoamihappyhacking/lazycat-skills
```

Once installed, your AI will automatically discover these skills. Try asking: "**Help me package this Docker project as a Lazycat LPK app**" or "**Set up a high-quality admin UI for my Lazycat project.**"

## 🤝 Contribution

We welcome Pull Requests from the community to improve documentation, add more automation scripts, or refine skill instructions!
