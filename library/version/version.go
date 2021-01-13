package version

import "fmt"

func ShowLogo(buildVersion, buildTime, commitID string) {
	//版本号
	//fmt.Println(" _   _  _____ _   _ _____  ___________ _____ \n| \\ | ||  ___| | | /  ___||  _  |  ___|_   _|\n|  \\| || |__ | | | \\ `--. | | | | |_    | |  \n| . ` ||  __|| | | |`--. \\| | | |  _|   | |  \n| |\\  || |___| |_| /\\__/ /\\ \\_/ / |     | |  \n\\_| \\_/\\____/ \\___/\\____/  \\___/\\_|     \\_/  ")
	fmt.Println("     _______.     ___       _______   ______     ______   \n    /       |    /   \\     /  _____| /  __  \\   /  __  \\  \n   |   (----`   /  ^  \\   |  |  __  |  |  |  | |  |  |  | \n    \\   \\      /  /_\\  \\  |  | |_ | |  |  |  | |  |  |  | \n.----)   |    /  _____  \\ |  |__| | |  `--'  | |  `--'  | \n|_______/    /__/     \\__\\ \\______|  \\______/   \\______/  \n                                                          ")
	fmt.Println("Version   ：", buildVersion)
	fmt.Println("BuildTime ：", buildTime)
	fmt.Println("CommitID  ：", commitID)
	fmt.Println("")

}
