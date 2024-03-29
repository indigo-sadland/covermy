<p align="center">
  <img width="400" alt="covermy"  src="https://github.com/indigo-sadland/covermy/blob/main/assets/logo.png?raw=true">
 </p>
 <p align="center"><em>Helper in storing commands and its' output!</em></p>

## Overview
`covermy` is a simple tool dedicated to help you stay organised. It relays on [Joplin](https://github.com/laurent22/joplin) - open source note-taking app, and uses Joplin API to create hierarchical notes to store bash commands along with its' output.

## Goal
I've made this tool with the intention to automate my bug bounty/pentest routine. I also hope that `covermy` will help not only me and not only bug bounty hunters or pentesters.

## Instalation
### Requirements
* Go 1.19+
  ```
  go install github.com/indigo-sadland/covermy@latest
  ```
* Joplin Desktop App

After installing Joplin you need to obtain API key. To do this, start Joplin and go to `Tools` -> `Options` -> `Web Clipper`. Copy the key and place it in `$USER_HOME/.local/share/covermy/api` as plain text.

## Usage
Simply add `covermy` before desired command.

### Demo

![covermy](https://user-images.githubusercontent.com/37074372/205365950-b72ebdf9-1cd9-45f7-b190-658a9ab6d540.gif)




 
