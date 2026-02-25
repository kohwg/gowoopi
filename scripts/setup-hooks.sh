#!/bin/sh
# Setup git hooks for this repository

git config core.hooksPath .githooks
chmod +x .githooks/*

echo "✅ Git hooks configured successfully!"
echo "   Commit messages will now be validated against conventional commits format."
