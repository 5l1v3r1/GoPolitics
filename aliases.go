package main

var BernieAliases = []string{
	"berniesanders",
	"bernie",
	"bernie sanders",
	"sanders",
}

var DonaldAliases = []string{
	"realdonaldtrump",
	"donaldtrump",
	"doanldjtrump",
	"donald trump",
	"donald j trump",
	"trump",
}

var HillaryAliases = []string{
	"clinton",
	"hillary",
	"hillaryclinton",
	"hillary clinton",
}

var TedAliases = []string{
	"cruz",
	"sentedcruz",
	"tedcruz",
	"ted cruz",
}

var AllAliases = append(append(BernieAliases, DonaldAliases...), append(HillaryAliases, TedAliases...)...)
