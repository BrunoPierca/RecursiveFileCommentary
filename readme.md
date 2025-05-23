# TypeScript NoCheck Adder

A simple Go utility that recursively adds `// @ts-nocheck` comments to all TypeScript files in a directory tree.


**⚠️ Important**: Always commit your changes to version control before running this tool, especially on important projects. While the tool is designed to be safe, having a backup is always recommended.

## Overview

This tool helps you quickly disable TypeScript type checking across all `.ts` and `.tsx` files in your project by adding the `// @ts-nocheck` directive. This is particularly useful when:

- Working with legacy TypeScript code that has many type errors
- Temporarily disabling type checking during major refactoring
- Converting JavaScript projects to TypeScript incrementally
- Dealing with third-party code that has type issues

## Features

- **Recursive Processing**: Scans all subdirectories automatically
- **Smart Placement**: Intelligently places `// @ts-nocheck` after file headers and initial comments
- **Duplicate Prevention**: Skips files that already have the directive
- **Detailed Reporting**: Shows exactly which files were processed
- **Error Handling**: Continues processing even if some files fail
- **Cross-Platform**: Works on Windows, macOS, and Linux

## Installation

### Prerequisites
- Go 1.16 or later installed on your system

### Option 1: Run Directly
1. Save the script as `add-ts-nocheck.go`
2. Navigate to your target directory
3. Run: `go run add-ts-nocheck.go`

### Option 2: Compile and Run
1. Compile: `go build add-ts-nocheck.go`
2. Run: `./add-ts-nocheck` (or `add-ts-nocheck.exe` on Windows)

## Usage

```bash
# Navigate to your project directory
cd /path/to/your/project

# Run the tool
go run add-ts-nocheck.go
```

### Example Output

```
Processing TypeScript files in: /Users/username/myproject
Added @ts-nocheck to: component.tsx
Added @ts-nocheck to: utils.ts
Skipping helper.ts (already has @ts-nocheck)
Added @ts-nocheck to: types.ts

=== Summary ===
Successfully processed: 3 files

Processed files:
  ✓ src/component.tsx
  ✓ src/utils.ts
  ✓ src/types.ts
```

## How It Works

The tool performs the following steps:

1. **Discovery**: Recursively walks through all directories starting from the current location
2. **Filtering**: Identifies files with `.ts` or `.tsx` extensions
3. **Analysis**: Checks if `// @ts-nocheck` already exists in each file
4. **Insertion**: Adds the directive at the appropriate location:
   - After any shebang lines
   - After initial comments or empty lines
   - Before the actual TypeScript code
5. **Reporting**: Provides detailed feedback on processed files

## File Structure Impact

### Before
```typescript
import React from 'react';

interface Props {
  name: string;
}

const Component: React.FC<Props> = ({ name }) => {
  return <div>Hello {name}</div>;
};
```

### After
```typescript
// @ts-nocheck
import React from 'react';

interface Props {
  name: string;
}

const Component: React.FC<Props> = ({ name }) => {
  return <div>Hello {name}</div>;
};
```

## Safety Features

- **Non-destructive**: Only adds comments, doesn't modify existing code
- **Idempotent**: Running multiple times won't create duplicate directives
- **Backup recommended**: While safe, it's good practice to commit your changes to version control first

## Limitations

- Only processes `.ts` and `.tsx` files
- Requires Go runtime or compilation
- Modifies files in place (no undo feature)

## Contributing

Feel free to submit issues or pull requests to improve this tool. Some potential enhancements:

- Add support for removing `@ts-nocheck` directives
- Include/exclude patterns for file filtering
- Backup file creation option
- Configuration file support

## License

This project is released under the MIT License. Feel free to use, modify, and distribute as needed.

## Troubleshooting

### Permission Errors
If you encounter permission errors, ensure you have write access to all files in the target directory.

### Go Not Found
Make sure Go is installed and available in your system PATH. Visit [golang.org](https://golang.org/dl/) for installation instructions.

### Large Projects
For very large projects (thousands of files), the tool may take a few moments to complete. Progress is shown in real-time.

---

**⚠️ Important**: Always commit your changes to version control before running this tool, especially on important projects. While the tool is designed to be safe, having a backup is always recommended.