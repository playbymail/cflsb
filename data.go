// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package sb

var (
	Households = map[FAMILY]Family{
		FAMILY_COOPER:   {Name: "Cooper", Crest: "Gold Crowns", Avatar: "crest_01.png"},
		FAMILY_WALKER:   {Name: "Walker", Crest: "Blue Fish", Avatar: "crest_06.png"},
		FAMILY_PAYNE:    {Name: "Payne", Crest: "Purple Grapes", Avatar: "crest_05.png"},
		FAMILY_HUGHES:   {Name: "Hughes", Crest: "Red Lions", Avatar: "crest_03.png"},
		FAMILY_FLETCHER: {Name: "Fletcher", Crest: "Black Swords", Avatar: "crest_02.png"},
		FAMILY_NASH:     {Name: "Nash", Crest: "Green Trees", Avatar: "crest_04.png"},
	}

	ActorDeck = []Actor{
		{Name: "Agincourt Duckbush", Benefactor: FAMILY_COOPER, Rank: 1, Page: 1, Row: 1, Col: 1},
		{Name: "Meese Markland", Benefactor: FAMILY_COOPER, Rank: 2, Page: 1, Row: 1, Col: 2},
		{Name: "Jupiter Browncakes", Benefactor: FAMILY_COOPER, Rank: 3, Page: 1, Row: 1, Col: 3},
		{Name: "Felicia Clover", Benefactor: FAMILY_COOPER, Rank: 4, Page: 1, Row: 2, Col: 1},
		{Name: "Lutward Stillman", Benefactor: FAMILY_COOPER, Rank: 5, Page: 1, Row: 2, Col: 2},
		{Name: "Tallboys Corklet", Benefactor: FAMILY_COOPER, Rank: 6, Page: 1, Row: 2, Col: 3},
		{Name: "Katherine Sarpow", Benefactor: FAMILY_COOPER, Rank: 7, Page: 1, Row: 3, Col: 1},
		{Name: "Brendan Silk", Benefactor: FAMILY_COOPER, Rank: 8, Page: 1, Row: 3, Col: 2},
		{Name: "Constanze Cooper", Benefactor: FAMILY_COOPER, Rank: 9, Page: 1, Row: 3, Col: 3},

		{Name: "Alana Petrie", Benefactor: FAMILY_WALKER, Rank: 1, Page: 2, Row: 1, Col: 1},
		{Name: "Duncan Pride", Benefactor: FAMILY_WALKER, Rank: 2, Page: 2, Row: 1, Col: 2},
		{Name: "Harriet Grove", Benefactor: FAMILY_WALKER, Rank: 3, Page: 2, Row: 1, Col: 3},
		{Name: "Quentin Welles", Benefactor: FAMILY_WALKER, Rank: 4, Page: 2, Row: 2, Col: 1},
		{Name: "Rosalind Brusque", Benefactor: FAMILY_WALKER, Rank: 5, Page: 2, Row: 2, Col: 2},
		{Name: "Urmond Wha", Benefactor: FAMILY_WALKER, Rank: 6, Page: 2, Row: 2, Col: 3},
		{Name: "Gontarde Goodwine", Benefactor: FAMILY_WALKER, Rank: 7, Page: 2, Row: 3, Col: 1},
		{Name: "Foucauld Stacks", Benefactor: FAMILY_WALKER, Rank: 8, Page: 2, Row: 3, Col: 2},
		{Name: "Beatrice Walker", Benefactor: FAMILY_WALKER, Rank: 9, Page: 2, Row: 3, Col: 3},

		{Name: "Salute Givenshy", Benefactor: FAMILY_PAYNE, Rank: 1, Page: 3, Row: 1, Col: 1},
		{Name: "Helmward Crouch", Benefactor: FAMILY_PAYNE, Rank: 2, Page: 3, Row: 1, Col: 2},
		{Name: "Ingarde Moonbeam", Benefactor: FAMILY_PAYNE, Rank: 3, Page: 3, Row: 1, Col: 3},
		{Name: "Stretch Coverlette", Benefactor: FAMILY_PAYNE, Rank: 4, Page: 3, Row: 2, Col: 1},
		{Name: "Lanie Foster", Benefactor: FAMILY_PAYNE, Rank: 5, Page: 3, Row: 2, Col: 2},
		{Name: "Hardwick McGill", Benefactor: FAMILY_PAYNE, Rank: 6, Page: 3, Row: 2, Col: 3},
		{Name: "Diamond Moreau", Benefactor: FAMILY_PAYNE, Rank: 7, Page: 3, Row: 3, Col: 1},
		{Name: "Victoria Flowers", Benefactor: FAMILY_PAYNE, Rank: 8, Page: 3, Row: 3, Col: 2},
		{Name: "Adalmar Payne", Benefactor: FAMILY_PAYNE, Rank: 9, Page: 3, Row: 3, Col: 3},

		{Name: "Ernesto Pleasir", Benefactor: FAMILY_HUGHES, Rank: 1, Page: 4, Row: 1, Col: 1},
		{Name: "Georges DeMaine", Benefactor: FAMILY_HUGHES, Rank: 2, Page: 4, Row: 1, Col: 2},
		{Name: "Tracy DesLions", Benefactor: FAMILY_HUGHES, Rank: 3, Page: 4, Row: 1, Col: 3},
		{Name: "Plonk Dargon", Benefactor: FAMILY_HUGHES, Rank: 4, Page: 4, Row: 2, Col: 1},
		{Name: "Carabelle Wildacre", Benefactor: FAMILY_HUGHES, Rank: 5, Page: 4, Row: 2, Col: 2},
		{Name: "Saxon Compote", Benefactor: FAMILY_HUGHES, Rank: 6, Page: 4, Row: 2, Col: 3},
		{Name: "Limerik Tealeaf", Benefactor: FAMILY_HUGHES, Rank: 7, Page: 4, Row: 3, Col: 1},
		{Name: "Bela Lightheart", Benefactor: FAMILY_HUGHES, Rank: 8, Page: 4, Row: 3, Col: 2},
		{Name: "Lohan Hughes", Benefactor: FAMILY_HUGHES, Rank: 9, Page: 4, Row: 3, Col: 3},

		{Name: "Grace York", Benefactor: FAMILY_FLETCHER, Rank: 1, Page: 5, Row: 1, Col: 1},
		{Name: "Wolfbert Grouse", Benefactor: FAMILY_FLETCHER, Rank: 2, Page: 5, Row: 1, Col: 2},
		{Name: "Cormac The Blade", Benefactor: FAMILY_FLETCHER, Rank: 3, Page: 5, Row: 1, Col: 3},
		{Name: "Arcar Portmanteau", Benefactor: FAMILY_FLETCHER, Rank: 4, Page: 5, Row: 2, Col: 1},
		{Name: "Vermilia Leagues", Benefactor: FAMILY_FLETCHER, Rank: 5, Page: 5, Row: 2, Col: 2},
		{Name: "Ermentrude Botas", Benefactor: FAMILY_FLETCHER, Rank: 6, Page: 5, Row: 2, Col: 3},
		{Name: "Alsous Grief", Benefactor: FAMILY_FLETCHER, Rank: 7, Page: 5, Row: 3, Col: 1},
		{Name: "Petunia Libet", Benefactor: FAMILY_FLETCHER, Rank: 8, Page: 5, Row: 3, Col: 2},
		{Name: "Vaughn Fletcher", Benefactor: FAMILY_FLETCHER, Rank: 9, Page: 5, Row: 3, Col: 3},

		{Name: "Hambard Brewer", Benefactor: FAMILY_NASH, Rank: 1, Page: 6, Row: 1, Col: 1},
		{Name: "Nanne Gambader", Benefactor: FAMILY_NASH, Rank: 2, Page: 6, Row: 1, Col: 2},
		{Name: "Evan Moore", Benefactor: FAMILY_NASH, Rank: 3, Page: 6, Row: 1, Col: 3},
		{Name: "Bella Briquote", Benefactor: FAMILY_NASH, Rank: 4, Page: 6, Row: 2, Col: 1},
		{Name: "Leslie Frye", Benefactor: FAMILY_NASH, Rank: 5, Page: 6, Row: 2, Col: 2},
		{Name: "Aitilde Flaire", Benefactor: FAMILY_NASH, Rank: 6, Page: 6, Row: 2, Col: 3},
		{Name: "Merelide VamPoore", Benefactor: FAMILY_NASH, Rank: 7, Page: 6, Row: 3, Col: 1},
		{Name: "Zachary Forester", Benefactor: FAMILY_NASH, Rank: 8, Page: 6, Row: 3, Col: 2},
		{Name: "Aelis Nash", Benefactor: FAMILY_NASH, Rank: 9, Page: 6, Row: 3, Col: 3},
	}

	PlayDeck = []Play{
		// page 1, row 1
		{Name: "Our Most Various Misdeeds", TheatricalValue: 5, Category: COMEDY, PointValue: 1, Page: 1, Row: 1, Col: 1},
		{Name: "The Jester", TheatricalValue: 5, Category: TRAGEDY, PointValue: 1, Page: 1, Row: 1, Col: 2},
		{Name: "The Fall of Ethelred", TheatricalValue: 6, Category: HISTORY, PointValue: 1, Page: 1, Row: 1, Col: 3, Benefactor: FAMILY_WALKER},
		// page 1, row 2
		{Name: "A Cobbler of Leeds", TheatricalValue: 6, Category: COMEDY, PointValue: 1, Page: 1, Row: 2, Col: 1},
		{Name: "The Duke's Demise", TheatricalValue: 6, Category: TRAGEDY, PointValue: 1, Page: 1, Row: 2, Col: 2},
		{Name: "A Most Impractical Hat", TheatricalValue: 7, Category: COMEDY, PointValue: 1, Page: 1, Row: 2, Col: 3},
		// page 1, row 3
		{Name: "Hastings", TheatricalValue: 8, Category: HISTORY, PointValue: 1, Page: 1, Row: 3, Col: 1, Benefactor: FAMILY_HUGHES},
		{Name: "The Wretched Wives of Washfield", TheatricalValue: 8, Category: TRAGEDY, PointValue: 1, Page: 1, Row: 3, Col: 2},
		{Name: "The Trials of Sisyphus", TheatricalValue: 9, Category: HISTORY, PointValue: 2, Page: 1, Row: 3, Col: 3},
		// page 2, row 1
		{Name: "The Lustful Lady Groutfoot", TheatricalValue: 9, Category: TRAGEDY, PointValue: 2, Page: 2, Row: 1, Col: 1},
		{Name: "Fizzleworth and Gumbody", TheatricalValue: 10, Category: COMEDY, PointValue: 2, Page: 2, Row: 1, Col: 2},
		{Name: "Robin and The Prince", TheatricalValue: 10, Category: HISTORY, PointValue: 2, Page: 2, Row: 1, Col: 3},
		// page 2, row 2
		{Name: "A Dance of the Steppes", TheatricalValue: 11, Category: COMEDY, PointValue: 2, Page: 2, Row: 2, Col: 1, Benefactor: FAMILY_FLETCHER},
		{Name: "Tiberius Tumbles", TheatricalValue: 11, Category: TRAGEDY, PointValue: 2, Page: 2, Row: 2, Col: 2, Benefactor: FAMILY_NASH},
		{Name: "The Black Prince", TheatricalValue: 12, Category: HISTORY, PointValue: 2, Page: 2, Row: 2, Col: 3, Benefactor: FAMILY_COOPER},
		// page 2, row 3
		{Name: "A Very Naughty Friar", TheatricalValue: 12, Category: COMEDY, PointValue: 2, Page: 2, Row: 3, Col: 1, Benefactor: FAMILY_NASH},
		{Name: "A Miserable Shame", TheatricalValue: 13, Category: TRAGEDY, PointValue: 3, Page: 2, Row: 3, Col: 2},
		{Name: "Polly's Peculiar Pranks", TheatricalValue: 13, Category: COMEDY, PointValue: 3, Page: 2, Row: 3, Col: 3},
		// page 3, row 1
		{Name: "The War of the Lancasters", TheatricalValue: 14, Category: HISTORY, PointValue: 3, Page: 3, Row: 1, Col: 1},
		{Name: "Gloria's Gambit", TheatricalValue: 14, Category: TRAGEDY, PointValue: 3, Page: 3, Row: 1, Col: 2},
		{Name: "The Triumph (of Queen Boudica)", TheatricalValue: 15, Category: HISTORY, PointValue: 3, Page: 3, Row: 1, Col: 3, Benefactor: FAMILY_PAYNE},
		// page 3, row 2
		{Name: "As With Anything Similar", TheatricalValue: 15, Category: TRAGEDY, PointValue: 3, Page: 3, Row: 2, Col: 1, Benefactor: FAMILY_HUGHES},
		{Name: "An Apple for LIttle Frederick", TheatricalValue: 16, Category: COMEDY, PointValue: 3, Page: 3, Row: 2, Col: 2, Benefactor: FAMILY_FLETCHER},
		{Name: "Empress Maud", TheatricalValue: 16, Category: HISTORY, PointValue: 3, Page: 3, Row: 2, Col: 3, Benefactor: FAMILY_WALKER},
		// page 3, row 3
		{Name: "The Goats of Whimsy Town", TheatricalValue: 17, Category: COMEDY, PointValue: 4, Page: 3, Row: 3, Col: 1},
		{Name: "Leopold's Lament", TheatricalValue: 17, Category: TRAGEDY, PointValue: 4, Page: 3, Row: 3, Col: 2},
		{Name: "Alfred the Great", TheatricalValue: 18, Category: HISTORY, PointValue: 4, Page: 3, Row: 3, Col: 3, Benefactor: FAMILY_NASH},
		// page 4, row 1
		{Name: "Magical Mirth", TheatricalValue: 18, Category: COMEDY, PointValue: 4, Page: 4, Row: 1, Col: 1, Benefactor: FAMILY_WALKER},
		{Name: "Prattle of Piteous Portent", TheatricalValue: 19, Category: TRAGEDY, PointValue: 4, Page: 4, Row: 1, Col: 2, Benefactor: FAMILY_COOPER},
		{Name: "Bottoms and Blossoms", TheatricalValue: 19, Category: COMEDY, PointValue: 4, Page: 4, Row: 1, Col: 3, Benefactor: FAMILY_HUGHES},
		// page 4, row 2
		{Name: "The Norman Crown", TheatricalValue: 20, Category: HISTORY, PointValue: 4, Page: 4, Row: 2, Col: 1, Benefactor: FAMILY_PAYNE},
		{Name: "The Wily Witches of Wackerfield", TheatricalValue: 20, Category: TRAGEDY, PointValue: 4, Page: 4, Row: 2, Col: 2, Benefactor: FAMILY_FLETCHER},
	}
)
