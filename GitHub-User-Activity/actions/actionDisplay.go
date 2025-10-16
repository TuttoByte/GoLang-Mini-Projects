package actions

import "fmt"

func DisplayGitHubActivities(activitiesList []ActivitiesInfoJSON) {

	for i := range activitiesList {

		switch activitiesList[i].Type {

		case "PushEvent":
			{
				var commitsNumber = len(activitiesList[i].Payload.Commits)
				fmt.Printf("Pushed %d commits to %s\n", commitsNumber, activitiesList[i].Repo.Name)
			}

		case "CreateEvent":
			{
				fmt.Printf("Created %s, named %s\n", activitiesList[i].Payload.Ref_type, activitiesList[i].Repo.Name)
			}

		case "IssuesEvent":
			{
				fmt.Printf("%s a new issue in %s\n", activitiesList[i].Payload.Action, activitiesList[i].Repo.Name)
			}

		case "WatchEvent":
			{
				fmt.Printf("Started watching %s\n", activitiesList[i].Repo.Name)
			}
		case "ForkEvent":
			{
				fmt.Printf("Forked %s\n", activitiesList[i].Repo.Name)
			}

		default:
			fmt.Printf("%s in %s\n", activitiesList[i].Type)
		}

	}
}
