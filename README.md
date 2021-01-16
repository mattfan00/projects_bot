# Projects Tracker 

A slack bot that tracks the projects the tech@nyu community makes 

## Routes 
`/slack`
- Used to handle the mention to the Projects Tracker bot 
- Also used to handle the Slack challenge

`/slack/all-projects`
- Displays a list of all of the projects 

`/slack/add-projects`
- Displays a modal that includes a form the user can fill out to add a new project
- Modal calls `/slack/interactive` to create a project with the details from the form 

`/slack/delete-projects`
- Displays buttons that the user can click to delete a project 
- The button that the user clicks is sent to `/slack/interactive` to delete that project

`/slack/interactive`
- Handles any of the interactions (button click, form submit, etc)
