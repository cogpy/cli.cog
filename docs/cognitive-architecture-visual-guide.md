# Cognitive Architecture Visual Guide

This document provides visual representations of key concepts from the cognitive architecture analysis.

---

## The Four Ways of Knowing

```
┌─────────────────────────────────────────────────────────────────┐
│                    INTEGRATED COGNITION                          │
│                                                                   │
│  ┌───────────────┐  ┌───────────────┐  ┌───────────────┐       │
│  │ PROPOSITIONAL │  │  PROCEDURAL   │  │ PERSPECTIVAL  │       │
│  │  (Knowing-    │  │  (Knowing-    │  │  (Knowing-    │       │
│  │    That)      │  │    How)       │  │    As)        │       │
│  ├───────────────┤  ├───────────────┤  ├───────────────┤       │
│  │ Facts         │  │ Skills        │  │ Frames        │       │
│  │ Beliefs       │  │ Abilities     │  │ Salience      │       │
│  │ Theory        │  │ Expertise     │  │ Aspect        │       │
│  │               │  │               │  │ Perception    │       │
│  │ Example:      │  │ Example:      │  │               │       │
│  │ "Paris is     │  │ Riding a      │  │ Example:      │       │
│  │  capital      │  │  bicycle      │  │ Duck-rabbit   │       │
│  │  of France"   │  │               │  │  gestalt      │       │
│  └───────────────┘  └───────────────┘  └───────────────┘       │
│                                                                   │
│                    ┌───────────────┐                            │
│                    │ PARTICIPATORY │                            │
│                    │  (Knowing-    │                            │
│                    │   By-Being)   │                            │
│                    ├───────────────┤                            │
│                    │ Identity      │                            │
│                    │ Transform.    │                            │
│                    │ Belonging     │                            │
│                    │ Co-identify   │                            │
│                    │               │                            │
│                    │ Example:      │                            │
│                    │ Being a       │                            │
│                    │  parent       │                            │
│                    └───────────────┘                            │
│                                                                   │
└─────────────────────────────────────────────────────────────────┘

Status in cli.cog:
  ✓ Propositional: Good (documentation, help text)
  ✓ Procedural: Moderate (command workflows)
  ⚠ Perspectival: Weak (no frame switching)
  ✗ Participatory: Absent (no transformation)
```

---

## Relevance Realization Process

```
┌──────────────────────────────────────────────────────────────────┐
│                  RELEVANCE REALIZATION CYCLE                      │
│                                                                    │
│   ∞ Infinite Possibilities                                        │
│        │                                                           │
│        ▼                                                           │
│   ┌─────────────┐                                                │
│   │   FILTER    │  Reduce to tractable set                       │
│   │             │  (1000s → 10s)                                 │
│   └──────┬──────┘                                                │
│          │                                                         │
│          ▼                                                         │
│   ┌─────────────┐                                                │
│   │   FRAME     │  Apply perspectival framing                    │
│   │             │  What becomes salient?                         │
│   └──────┬──────┘                                                │
│          │                                                         │
│          ▼                                                         │
│   ┌─────────────┐                                                │
│   │  SALIENCE   │  Map salience landscape                        │
│   │             │  What stands out?                              │
│   └──────┬──────┘                                                │
│          │                                                         │
│          ▼                                                         │
│   ┌─────────────┐                                                │
│   │  OPPONENT   │  Balance tradeoffs:                            │
│   │ PROCESSING  │  - Explore ↔ Exploit                          │
│   │             │  - Speed ↔ Accuracy                           │
│   └──────┬──────┘  - Certain ↔ Open                             │
│          │                                                         │
│          ▼                                                         │
│   ┌─────────────┐                                                │
│   │    RANK     │  Order by relevance                            │
│   │             │                                                 │
│   └──────┬──────┘                                                │
│          │                                                         │
│          ▼                                                         │
│   ┌─────────────┐                                                │
│   │   SELECT    │  Choose action(s)                              │
│   │             │                                                 │
│   └──────┬──────┘                                                │
│          │                                                         │
│          ▼                                                         │
│   ┌─────────────┐                                                │
│   │  FEEDBACK   │  Learn from outcome                            │
│   │   LOOP      │  Update relevance criteria                     │
│   └──────┬──────┘                                                │
│          │                                                         │
│          └──────────┐                                            │
│                      ▼                                            │
│              (Feed Back to Filter)                                │
│                                                                    │
│   Status in cli.cog:                                              │
│   ✗ Currently: User does all relevance realization               │
│   ○ Proposed: System assists in determining relevance            │
│                                                                    │
└──────────────────────────────────────────────────────────────────┘
```

---

## 4E Cognition Framework

```
┌──────────────────────────────────────────────────────────────────┐
│                      4E COGNITION                                 │
│                                                                    │
│  ┌───────────────────┐           ┌───────────────────┐          │
│  │    EMBODIED       │           │     EMBEDDED      │          │
│  │   (Sensorimotor)  │           │  (Environmental)  │          │
│  ├───────────────────┤           ├───────────────────┤          │
│  │ Body shapes mind  │           │ Context-dependent │          │
│  │ Action-perception │           │ Affordances       │          │
│  │ Somatic intel.    │           │ Niche construction│          │
│  │                   │           │                   │          │
│  │ Status: ⚠ Weak   │           │ Status: ✓ Moderate│          │
│  │ - Text only       │           │ - Good GitHub     │          │
│  │ - No sensorimotor │           │   integration     │          │
│  └───────────────────┘           └───────────────────┘          │
│                                                                    │
│  ┌───────────────────┐           ┌───────────────────┐          │
│  │     ENACTED       │           │     EXTENDED      │          │
│  │ (Action-Based)    │           │   (Distributed)   │          │
│  ├───────────────────┤           ├───────────────────┤          │
│  │ Learning through  │           │ Beyond brain      │          │
│  │   interaction     │           │ Tool integration  │          │
│  │ Sensorimotor      │           │ Social cognition  │          │
│  │   contingencies   │           │ Cognitive offload │          │
│  │                   │           │                   │          │
│  │ Status: ⚠ Weak   │           │ Status: ✓ Strong │          │
│  │ - No learning     │           │ - Agent daemons   │          │
│  │ - No adaptation   │           │ - Constellations  │          │
│  └───────────────────┘           └───────────────────┘          │
│                                                                    │
└──────────────────────────────────────────────────────────────────┘
```

---

## Perspectival Frame Switching

```
┌──────────────────────────────────────────────────────────────────┐
│                    MULTIPLE FRAMES                                │
│                                                                    │
│        Same Situation, Different Salience:                        │
│                                                                    │
│  ┌────────────────┐  ┌────────────────┐  ┌────────────────┐    │
│  │   TECHNICAL    │  │  USER EXPERIENCE│  │ BUSINESS VALUE │    │
│  │     FRAME      │  │      FRAME      │  │     FRAME      │    │
│  └────────────────┘  └────────────────┘  └────────────────┘    │
│                                                                    │
│  Salient:          Salient:           Salient:                   │
│  • Code quality    • Usability        • ROI                      │
│  • Performance     • Joy              • Strategic fit             │
│  • Architecture    • Accessibility    • Cost                     │
│                                                                    │
│  Blindspot:        Blindspot:         Blindspot:                 │
│  • User needs      • Tech debt        • Code quality             │
│  • Team dynamics   • Performance      • Developer joy            │
│                                                                    │
│                                                                    │
│                    WISDOM MODE:                                   │
│                 Hold Multiple Frames                              │
│                   Simultaneously                                  │
│                                                                    │
│     ┌──────────────────────────────────────────┐                │
│     │                                            │                │
│     │  Technical ⟷ User ⟷ Business ⟷ Ethical │                │
│     │                                            │                │
│     │        Balanced Awareness Across          │                │
│     │          Multiple Perspectives            │                │
│     │                                            │                │
│     │           (Sophrosyne)                    │                │
│     │                                            │                │
│     └──────────────────────────────────────────┘                │
│                                                                    │
│  Status in cli.cog:                                               │
│  ✗ Currently: Single fixed frame                                 │
│  ○ Proposed: Dynamic frame switching                             │
│                                                                    │
└──────────────────────────────────────────────────────────────────┘
```

---

## Meta-Cognitive Monitoring

```
┌──────────────────────────────────────────────────────────────────┐
│              META-COGNITIVE MONITORING STACK                      │
│                                                                    │
│  Level 3: Meta-Meta-Cognition                                    │
│  ┌────────────────────────────────────────────────────────┐     │
│  │ "How is my monitoring itself?"                          │     │
│  │ Monitoring quality assessment                           │     │
│  └────────────────────────────────────────────────────────┘     │
│                            ▲                                       │
│                            │                                       │
│  Level 2: Meta-Cognition                                         │
│  ┌────────────────────────────────────────────────────────┐     │
│  │ "What am I thinking about my thinking?"                │     │
│  │ Pattern detection, blindspot identification            │     │
│  │ Stuck detection, strategy evaluation                   │     │
│  └────────────────────────────────────────────────────────┘     │
│                            ▲                                       │
│                            │                                       │
│  Level 1: First-Order Cognition                                  │
│  ┌────────────────────────────────────────────────────────┐     │
│  │ "What am I doing right now?"                           │     │
│  │ Current task, active strategy, progress               │     │
│  └────────────────────────────────────────────────────────┘     │
│                                                                    │
│                                                                    │
│  Common Patterns Detected:                                        │
│                                                                    │
│  😵 Spinning Wheels    - Repeating without progress              │
│  🔒 Frame Fixation     - Stuck in single perspective             │
│  ⏸️  Analysis Paralysis - Over-analyzing without action          │
│  🌊 Productive Flow    - Smooth progress (celebrate!)            │
│  🚀 Premature Optimize - Optimizing before understanding         │
│                                                                    │
│  Status in cli.cog:                                               │
│  ✗ Currently: Declared but not implemented                       │
│  ○ Proposed: Real recursive self-monitoring                      │
│                                                                    │
└──────────────────────────────────────────────────────────────────┘
```

---

## Transformation Through Practice

```
┌──────────────────────────────────────────────────────────────────┐
│              WISDOM CULTIVATION JOURNEY                           │
│                                                                    │
│  Time ────────────────────────────────────────────────▶         │
│                                                                    │
│   Novice          Competent         Proficient        Expert     │
│     │                │                  │                │        │
│     │                │                  │                │        │
│     │    ┌───────────▼──────────────────▼────────────────▼───┐  │
│     │    │                                                     │  │
│     │    │        PARTICIPATORY TRANSFORMATION                │  │
│     │    │                                                     │  │
│     │    │  • Identity shifts from tool-user to practitioner │  │
│     │    │  • Capacities develop (not just skills)          │  │
│     │    │  • Wisdom emerges through practice                │  │
│     │    │                                                     │  │
│     │    │  Capacity Development:                            │  │
│     │    │                                                     │  │
│     │    │  Relevance Realization:    ⭐⭐⭐⭐⭐              │  │
│     │    │  Perspectival Flexibility: ⭐⭐⭐⭐☆              │  │
│     │    │  Meta-Cognition:          ⭐⭐⭐⭐⭐              │  │
│     │    │  Sophrosyne:              ⭐⭐⭐⭐☆              │  │
│     │    │  Pattern Recognition:      ⭐⭐⭐⭐⭐              │  │
│     │    │                                                     │  │
│     │    └─────────────────────────────────────────────────────┘  │
│     │                                                              │
│     ▼                                                              │
│  Tool User ──────────────────────────────▶ Wisdom Practitioner   │
│                                                                    │
│  Practices That Transform:                                        │
│                                                                    │
│  1. Meta-Cognitive Debugging    (15 min, daily)                  │
│  2. Perspectival Flexibility    (10 min, 3x/week)                │
│  3. Relevance Calibration       (20 min, weekly)                 │
│  4. Sophrosyne Cultivation      (25 min, 2x/week)                │
│                                                                    │
│  Status in cli.cog:                                               │
│  ✗ Currently: No participatory knowing or transformation         │
│  ○ Proposed: Structured wisdom cultivation practices             │
│                                                                    │
└──────────────────────────────────────────────────────────────────┘
```

---

## Opponent Processing Balances

```
┌──────────────────────────────────────────────────────────────────┐
│                OPPONENT PROCESSING TRADEOFFS                      │
│                                                                    │
│  These balances must be continuously optimized:                   │
│                                                                    │
│  Exploration ←──────────────────────→ Exploitation               │
│  [====·====================]                                      │
│   Try new            Use what works                               │
│   approaches         well already                                 │
│                                                                    │
│  Speed ←────────────────────────────→ Accuracy                   │
│  [=============·===========]                                      │
│   Quick         Thorough and                                      │
│   response      careful                                           │
│                                                                    │
│  Certainty ←────────────────────────→ Openness                   │
│  [==========·==============]                                      │
│   Confidence    Open to                                           │
│   in answer     alternatives                                      │
│                                                                    │
│  Breadth ←──────────────────────────→ Depth                      │
│  [============·============]                                      │
│   Survey        Deep dive                                         │
│   many areas    single area                                       │
│                                                                    │
│                                                                    │
│  Context-Dependent Presets:                                       │
│                                                                    │
│  🚀 Prototyping:    High exploration, speed, openness, breadth   │
│  🏭 Production:     High exploitation, accuracy, certainty, depth│
│  🔬 Research:       High exploration, accuracy, openness, breadth │
│  🐛 Debugging:      Medium exploration, accuracy, certainty, depth│
│  📈 Optimization:   Low exploration, high accuracy, certainty     │
│                                                                    │
│  Status in cli.cog:                                               │
│  ✗ Currently: No explicit tradeoff management                    │
│  ○ Proposed: Configurable opponent processing                    │
│                                                                    │
└──────────────────────────────────────────────────────────────────┘
```

---

## System Architecture Evolution

```
┌──────────────────────────────────────────────────────────────────┐
│                     CURRENT STATE                                 │
│                                                                    │
│              User                                                 │
│                │                                                   │
│                ▼                                                   │
│           ┌────────┐                                              │
│           │   gh   │  ← Fixed commands                            │
│           └────────┘    Execute tasks                             │
│                │                                                   │
│                ▼                                                   │
│          GitHub API                                               │
│                                                                    │
│  Characteristics:                                                 │
│  • User realizes all relevance                                    │
│  • Fixed command structure                                        │
│  • No learning or adaptation                                      │
│  • Tool-based relationship                                        │
│                                                                    │
└──────────────────────────────────────────────────────────────────┘

                              ▼
                    TRANSFORMATION
                              ▼

┌──────────────────────────────────────────────────────────────────┐
│                    PROPOSED STATE                                 │
│                                                                    │
│              Developer/Practitioner                               │
│                │         ▲                                        │
│                ▼         │ Transformation                         │
│           ┌──────────────────────┐                               │
│           │   Cognitive Layer    │                               │
│           ├──────────────────────┤                               │
│           │ • Relevance Engine   │ ← Realizes relevance          │
│           │ • Frame Manager      │ ← Switches perspectives       │
│           │ • Meta-Monitor       │ ← Self-awareness              │
│           │ • Practice Framework │ ← Cultivates wisdom           │
│           └──────────────────────┘                               │
│                │         ▲                                        │
│                ▼         │ Feedback                               │
│           ┌──────────────────────┐                               │
│           │   Execution Layer    │                               │
│           ├──────────────────────┤                               │
│           │ • Commands           │                               │
│           │ • Agent Daemons      │                               │
│           │ • Constellations     │                               │
│           └──────────────────────┘                               │
│                │                                                   │
│                ▼                                                   │
│          GitHub Ecosystem                                         │
│                                                                    │
│  Characteristics:                                                 │
│  • System assists relevance realization                           │
│  • Dynamic, context-aware                                         │
│  • Learns and adapts                                              │
│  • Practice-based relationship                                    │
│  • Transforms practitioner                                        │
│                                                                    │
└──────────────────────────────────────────────────────────────────┘
```

---

## Quick Reference: Key Concepts

| Concept | Definition | Current State | Proposed |
|---------|------------|---------------|----------|
| **Relevance Realization** | Determining what matters from infinite possibilities | User does it | System assists |
| **Propositional Knowing** | Knowledge of facts (knowing-that) | ✓ Good | ✓ Maintain |
| **Procedural Knowing** | Knowledge of skills (knowing-how) | ✓ Moderate | ✓ Enhance |
| **Perspectival Knowing** | Knowledge of framing (knowing-as) | ⚠️ Weak | ○ Add frames |
| **Participatory Knowing** | Transformative knowing (knowing-by-being) | ✗ Absent | ○ Add practices |
| **Sophrosyne** | Balanced judgment, optimal self-regulation | ✗ Absent | ○ Add balance |
| **Meta-Cognition** | Thinking about thinking | ✗ Not impl. | ○ Implement |
| **4E Cognition** | Embodied, Embedded, Enacted, Extended | Mixed | ○ Complete |
| **Opponent Processing** | Balancing cognitive tradeoffs | ✗ Absent | ○ Add controls |
| **Wisdom Cultivation** | Systematic relevance realization optimization | ✗ Absent | ○ Add practices |

Legend: ✓ Present, ⚠️ Partial, ✗ Absent, ○ Proposed

---

## Further Reading

- Main analysis: `cognitive-architecture-analysis.md`
- Implementation roadmap: `cognitive-enhancements-roadmap.md`
- Executive summary: `cognitive-architecture-executive-summary.md`
