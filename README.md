portscanner/
│
├── cmd/
│   └── portscanner/
│       └── main.go          # Entry point of the application
│
├── internal/
│   ├── scanner/
│   │   ├── tcp.go           # Contains TCP scanning logic
│   │   ├── udp.go           # Contains UDP scanning logic
│   │   └── syn.go           # (Optional) SYN scanning logic, more advanced
│   │
│   └── ui/
│       └── tui.go           # Terminal UI logic using a package like tview
│
├── pkg/
│   └── utils/
│       └── utils.go         # Helper functions (e.g., parsing ports, error handling)
│
├── go.mod                   # Go module file for dependency management
├── go.sum                   # Go module dependency file (generated)
└── README.md                # Project description and instructions