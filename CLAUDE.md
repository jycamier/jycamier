# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Hugo static site generator project. The codebase follows Hugo's standard directory structure:

- `content/` - Markdown content files
- `layouts/` - HTML template files
- `static/` - Static assets (CSS, JS, images)
- `themes/` - Hugo themes
- `archetypes/` - Content templates
- `data/` - Data files (YAML, JSON, TOML)
- `i18n/` - Internationalization files
- `hugo.toml` - Hugo configuration file

## Development Commands

### Build and Development
- `hugo` - Build the static site (output goes to `public/` directory)
- `hugo server` or `hugo serve` - Start development server with live reload (typically on http://localhost:1313)
- `hugo server --drafts` - Include draft content in development server
- `hugo --minify` - Build with minified output for production

### Content Management  
- `hugo new posts/my-post.md` - Create new content using archetype template
- `hugo new site <name>` - Create new Hugo site (if starting from scratch)

### Utility Commands
- `hugo version` - Check Hugo version
- `hugo help` - Get help on Hugo commands
- `hugo list drafts` - List all draft content
- `hugo list future` - List content with future publish dates

## Site Configuration

The main configuration is in `hugo.toml`. Currently configured as a basic Hugo site with:
- Base URL: https://example.org/
- Language: en-us  
- Title: "My New Hugo Site"

## Content Structure

Content uses Hugo's front matter format with TOML syntax:
```toml
+++
date = '2023-01-01'
draft = true
title = 'Page Title'
+++
```

## Theme and Layout System

This project uses Hugo's template system. Templates in `layouts/` override theme templates. The directory structure follows Hugo's lookup order for templates.