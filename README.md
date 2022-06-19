<p align="center">
  <img src=".github/brand/icon.png" height="128">
</p>

<h1 align="center">
  Ethereal Git
</h1>
<h3 align="center">
	A modern and comfortable Git desktop client.
</h3>
<p align="center">
  <a href="https://github.com/StarlaneStudios/EtherealGit/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/StarlaneStudios/EtherealGit"> 
  </a>
  <img src="https://img.shields.io/discord/414532188722298881">
  <img src="https://img.shields.io/github/repo-size/StarlaneStudios/EtherealGit">
  <img src="https://img.shields.io/github/contributors/StarlaneStudios/EtherealGit">
</p>

<br>

## 🚧 Disclaimer 🚧
This project is currently in heavy development and is still lacking many features. Do not use in production!

## Introduction

Ethereal Git is a modern Git desktop client focused on productivity and ease of use.

## Live Development

To run in live development mode, run `wails dev` in the project directory. In another terminal, go into the `frontend`
directory and run `npm run dev`. The frontend dev server will run on http://localhost:34115. Connect to this in your
browser and connect to your application.

## Generate Go bindings

To generate the callable Go bindings for JavaScript, run `wails generate module`.

## Building

To build a redistributable, production mode package, use `wails build`.
