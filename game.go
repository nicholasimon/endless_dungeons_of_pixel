package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// MARK: var
var (

	//TEST
	testv1, testv2 rl.Vector2
	testangle      float32

	//HELP
	helpon bool

	//ENDGAME
	endgamewindow bool

	//AUDIO
	intromusic               rl.Music
	backmusic                []rl.Music
	levmusic                 rl.Music
	musicon, soundfxon, mute bool
	soundvol                 = float32(0.2)

	swingaud, zapaud, collectaud, springaud, digaud, teleportpickupaud, wateraud, spikesaud, cactusaud, playerdamageaud, gemaud, campaud, stairsaud, eataud, monsterhit1aud, monsterhit2aud, monsterhit3aud, monsterhit4aud, monsterhit5aud, monsterhit6aud, monsterhit7aud, chestaud, weedaud, bookcaseaud rl.Sound

	scrollaud []rl.Sound

	//DIED
	diedy, diedfade float32

	//INTRO
	introon, intro2on, intro3on, intro4on, introlr                      bool
	introx, introy, introx2, introy2, introfade, introfade2, introplayx float32
	introtxtx, introtxtx2, introtxty                                    int32
	introtimer                                                          int

	//SCORE & TIMER
	scoreon                       bool
	runtimer                      int
	runmin, runsecs, scoreontimer int

	//SHOP
	shopon, closeshop, nokeys                    bool
	shoprecs, shoprecsorig                       []rl.Rectangle
	shoprecx, shoprectotallen                    float32
	keystotal, gemstotal                         int
	shopkeyimg, shopgemimg                       rl.Rectangle
	shopitems                                    []xshopitem
	nokeystimer                                  int32
	randomarmoron, cactusimmunity, spikeimmunity bool

	//PETS
	pets     []xpet
	petnames = []string{"Rex", "Fido", "Marmalade", "Muffins", "Odie", "Goblin", "Monty", "Churchill", "Rudolph", "Blitzen", "Pikachu", "Fonzie", "Homer", "Cartman", "Urkel", "Phineas", "Mac Daddy", "Chuck Norris", "Marshmallow", "Pee Wee", "Chunk", "Chickpea", "Porkchop", "Cocoa Puff", "Tic Tac", "Zipper", "Mister Miyagi", "Groucho Barks", "Bark Twain", "Mary Puppins"}
	//ETC
	refillcost                            = 1
	pigeononoff, pigeonon, pigeonflyingon bool
	pigeonx, pigeony                      float32
	pigeontimer                           int32
	springswitch, springswitch2           bool
	springv2                              rl.Vector2

	pigeonwords = []string{"hello hello", "whose a pretty boy then?", "in my previous life I was a parrot on the jolly roger", "spread my wings and fly away", "I believe I can fly", "I am not related to chickens", "Do you have any of my feathers in your pillow?", "Look out the window, I think that might be my cousin", "Do you have any idea what a pigeon is doing in this game?", "I couldn't find work in any other games, so here we are", "I don't do impressions", "Getting back to nature", "You really should get out more", "Another dingy day in the dungeon", "Did you notice the other pigeon was a substitute? I was on lunch"}

	pigeonphrase      string
	chosepigeonphrase bool

	//INVEN
	destroyitemon, destroybeltitemon                                   bool
	invenrec, destroyrec, destroybeltrec                               rl.Rectangle
	destroytimer, destroybelttimer                                     int32
	destroynum, destroybeltnum                                         int
	inven, beltinven                                                   []xobj
	invencurrentnum, activammonum, activweaponnum, beltinvencurrentnum int

	//LEGENDARY QUEST ITEMS
	questitemon      bool
	questitemv2      rl.Vector2
	questitemroomnum int
	uplegendaryon    bool
	armorsetcount    = make([]int, 6)
	armorsetactive   bool

	//POTIONS
	activpotions []xobj
	potiontimer  = fps * 60

	//MAGIC
	magicon       bool
	magictimer    int32
	activscroll   = xobj{}
	activmagic    []xmagic
	frograintimer int32

	//SETTINGS
	settingson bool

	//BOSS
	boss        []xboss
	nextbossnum = rInt(4, 8)

	//MAP
	mapon, teleporton bool

	//WEAPONS
	weaponv2                          rl.Vector2
	activweapons                      []xobj
	weaponro1, weaponro2, weaponro3   bool
	weaponrotimer                     int32
	weaponrotime                      = int32(2)
	autoswitchweapons, autoswitchammo = true, true
	weaponrangeon                     bool
	weaponrangetimer                  int32

	//MONSTERS
	monsters            []xmonster
	vismonsters         []xmonster
	monsteranimatetimer = 12
	dedmonsters         []xdedmonster
	showmonshp          bool
	enemybullets        []xobj
	monsternumlevel     int
	clearedmonsters     bool
	clearedleveltimer   int32
	clearedlevellootnum int

	monsternames = []string{"Liam", "Noah", "Oliver", "Elijah", "James", "William", "Benjamin", "Lucas", "Henry", "Theodore", "Jack", "Levi", "Alexander", "Jackson", "Mateo", "Daniel", "Michael", "Mason", "Sebastian", "Ethan", "Logan", "Owen", "Samuel", "Jacob", "Asher", "Aiden", "John", "Joseph", "Wyatt", "David", "Leo", "Luke", "Julian", "Hudson", "Grayson", "Matthew", "Ezra", "Gabriel", "Carter", "Isaac", "Jayden", "Luca", "Anthony", "Dylan", "Lincoln", "Thomas", "Maverick", "Elias", "Josiah", "Charles", "Caleb", "Christopher", "Ezekiel", "Miles", "Jaxon", "Isaiah", "Andrew", "Joshua", "Nathan", "Nolan", "Adrian", "Cameron", "Santiago", "Eli", "Aaron", "Ryan", "Cooper", "Angel", "Waylon", "Easton", "Kai", "Christian", "Landon", "Colton", "Roman", "Axel", "Brooks", "Jonathan", "Robert", "Jameson", "Ian", "Everett", "Greyson", "Wesley", "Jeremiah", "Hunter", "Leonardo", "Jordan", "Jose", "Bennett", "Silas", "Nicholas", "Parker", "Beau", "Weston", "Connor", "Austin", "Carson", "Dominic", "Xavier", "Jaxson", "Jace", "Emmett", "Adam", "Declan", "Rowan", "Micah", "Kayden", "Gael", "River", "Ryder", "Kingston", "Damian", "Sawyer", "Luka", "Evan", "Vincent", "Legend", "Myles", "Harrison", "August", "Bryson", "Amir", "Giovanni", "Chase", "Diego", "Milo", "Jasper", "Walker", "Jason", "Brayden", "Cole", "Nathaniel", "George", "Lorenzo", "Zion", "Luis", "Archer", "Enzo", "Jonah", "Thiago", "Theo", "Ayden", "Zachary", "Calvin", "Braxton", "Ashton", "Rhett", "Atlas", "Jude", "Bentley", "Carlos", "Ryker", "Adriel", "Arthur", "Ace", "Tyler", "Jayce", "Max", "Elliot", "Graham", "Kaiden", "Maxwell", "Juan", "Dean", "Matteo", "Malachi", "Ivan", "Elliott", "Jesus", "Emiliano", "Messiah", "Gavin", "Maddox", "Camden", "Hayden", "Leon", "Antonio", "Justin", "Tucker", "Brandon", "Kevin", "Judah", "Finn", "King", "Brody", "Xander", "Nicolas", "Charlie", "Arlo", "Emmanuel", "Barrett", "Felix", "Alex", "Miguel", "Abel", "Alan", "Beckett", "Amari", "Karter", "Timothy", "Abraham", "Jesse", "Zayden", "Blake", "Alejandro", "Dawson", "Tristan", "Victor", "Avery", "Joel", "Grant", "Eric", "Patrick", "Peter", "Richard", "Edward", "Andres", "Emilio", "Colt", "Knox", "Beckham", "Adonis", "Kyrie", "Matias", "Oscar", "Lukas", "Marcus", "Hayes", "Caden", "Remington", "Griffin", "Nash", "Israel", "Steven", "Holden", "Rafael", "Zane", "Jeremy", "Kash", "Preston", "Kyler", "Jax", "Jett", "Kaleb", "Riley", "Simon", "Phoenix", "Javier", "Bryce"}

	//MSG
	msgs        []string
	newmsgtimer int32
	displaymsgs bool

	//UI
	statsrec, msgrec, footerrec rl.Rectangle

	//BACKG
	backgrec rl.Rectangle

	//FX
	fxname      string
	fxrec       rl.Rectangle
	drawfxon    bool
	scan, ghost bool
	scanlines   []xscanline
	fx          []xfx
	blood       []xcircle
	xplodecircs []xcircle

	//PLAYER
	player                                    = xplayer{}
	killcount, monsterkills, score, bosskills int

	animateplayer, playeremoteon, emoteselected, invinciblemode, switchinvincible, deathon, died bool

	weapon = xobj{}
	armor  = xobj{}
	jewel  = xobj{}
	belt   = xobj{}

	//LEVEL
	level                            []xroom
	visroom                          []xroom
	origlevellen                     int
	levwid, levheig                  float32
	levtl, levtr, levbl, levbr       rl.Vector2
	levboundrec                      rl.Rectangle
	currentlevelnum, maxlevelreached int
	changelevelon                    bool

	//TILES
	basetile = float32(16)
	multi    = float32(3)
	tilesize = basetile * multi

	//SCR
	scrwf32, scrhf32      float32
	scrw, scrh            int32
	scrhint, scrwint      int
	scrcnt                rl.Vector2
	borderrec, visiblerec rl.Rectangle
	defw                  = float32(1280)
	defh                  = float32(720)

	//TXT
	txts   = int32(10)
	txtm   = int32(20)
	txtl   = int32(30)
	txtl2  = int32(40)
	txtxl  = int32(80)
	txtdef = txtm

	//TIMERS
	fps                                  = int32(60)
	fpsint                               = int(fps)
	frames                               int
	fadeblink                            = float32(0.2)
	fadeblinkon                          bool
	playeranimatetimer, playeremotetimer int
	pause                                bool

	onoff1, onoff2, onoff3, onoff6, onoff10, onoff15, onoff30, onoff60 bool
	//CAMS
	camera, cammap rl.Camera2D

	//INP
	mousev2, mousev2world rl.Vector2
	selpoint              rl.Vector2
	selrec                rl.Rectangle
	selroom               int
	mouseclicknum         int
	clickpause            int32
	mousroomnum           int

	//DEV
	dev, dev2, teston bool

	//BLANKS
	blankv2  = rl.NewVector2(77777777777777777, 77777777777777777)
	blankint = 77777777777777777

	//IMGS
	imgs   rl.Texture2D
	origin = rl.NewVector2(0, 0)

	//MARK: PLAYER IMGS PLAYER IMGS PLAYER IMGS PLAYER IMGS PLAYER IMGS
	//player
	playerimg        = rl.NewRectangle(1, 92, 16, 16)
	playerlimg       = rl.NewRectangle(66, 110, 16, 16)
	poisonedimg      = rl.NewRectangle(133, 320, 16, 16)
	fireresistimg    = rl.NewRectangle(759, 249, 16, 16)
	poisonresistimg  = rl.NewRectangle(743, 249, 16, 16)
	sickimg          = rl.NewRectangle(778, 249, 16, 16)
	diseaseimmuneimg = rl.NewRectangle(195, 322, 16, 16)
	intimg           = rl.NewRectangle(4, 47, 16, 16)
	hpmaximg         = rl.NewRectangle(26, 47, 16, 16)
	strimg           = rl.NewRectangle(48, 46, 18, 18)
	deximg           = rl.NewRectangle(72, 48, 16, 16)
	lukimg           = rl.NewRectangle(92, 46, 18, 18)

	//boss
	mushroombossl []rl.Rectangle
	mushroomboss  []rl.Rectangle
	dinoboss      []rl.Rectangle
	dinobossl     []rl.Rectangle
	skullboss     []rl.Rectangle
	skullbossl    []rl.Rectangle
	radishboss    []rl.Rectangle
	radishbossl   []rl.Rectangle
	spikeboss     []rl.Rectangle
	spikebossl    []rl.Rectangle
	ghostboss     []rl.Rectangle
	ghostbossl    []rl.Rectangle
	reaperboss    []rl.Rectangle
	reaperbossl   []rl.Rectangle
	orcboss       []rl.Rectangle
	orcbossl      []rl.Rectangle
	slimeboss     []rl.Rectangle
	slimebossl    []rl.Rectangle

	//etc
	audioicon       = rl.NewRectangle(209, 77, 14, 14)
	teddyimg        = rl.NewRectangle(475, 296, 18, 18)
	hpimg           = rl.NewRectangle(0, 76, 16, 16)
	flameimg        = rl.NewRectangle(111, 318, 18, 18)
	meteorimg       = rl.NewRectangle(2, 340, 32, 32)
	newmsgimg       = rl.NewRectangle(113, 77, 12, 12)
	fireballimg     = rl.NewRectangle(1058, 196, 14, 14)
	tickimg         = rl.NewRectangle(127, 74, 17, 17)
	cancelimg       = rl.NewRectangle(147, 76, 14, 14)
	coinimg         = rl.NewRectangle(1, 133, 16, 16)
	downarrow2img   = rl.NewRectangle(168, 76, 14, 14)
	uparrow2img     = rl.NewRectangle(186, 77, 14, 14)
	frisbeeimg      = rl.NewRectangle(153, 321, 16, 16)
	frogimg         = rl.NewRectangle(1303, 15, 64, 64)
	pigeonimg       = rl.NewRectangle(8, 488, 16, 16)
	pigeonflyingimg = rl.NewRectangle(57, 506, 16, 16)
	springimg       = rl.NewRectangle(43, 352, 16, 16)
	spikeimg        = rl.NewRectangle(97, 357, 13, 13)
	lockimg         = rl.NewRectangle(219, 323, 14, 14)
	soldimg         = rl.NewRectangle(241, 322, 16, 16)

	teleportimg   = rl.NewRectangle(173, 321, 16, 16)
	teleportimgro float32

	emoteimg    rl.Rectangle
	emoteimgx   float32
	emoteimgnum int
	emoteimgs   []rl.Rectangle

	//intro
	gologo     = rl.NewRectangle(3, 897, 221, 300)
	rayliblogo = rl.NewRectangle(275, 926, 256, 256)

	//pets
	mouse1iimg = rl.NewRectangle(145, 487, 16, 16)
	mouse1rimg = rl.NewRectangle(81, 487, 16, 16)
	mouse1limg = rl.NewRectangle(129, 505, 16, 16)

	mouse2iimg = rl.NewRectangle(76, 527, 16, 16)
	mouse2rimg = rl.NewRectangle(10, 527, 16, 16)
	mouse2limg = rl.NewRectangle(59, 546, 16, 16)

	sheepiimg = rl.NewRectangle(79, 569, 16, 16)
	sheeprimg = rl.NewRectangle(11, 569, 16, 16)
	sheeplimg = rl.NewRectangle(58, 589, 16, 16)

	dog1iimg = rl.NewRectangle(78, 613, 16, 16)
	dog1rimg = rl.NewRectangle(9, 613, 16, 16)
	dog1limg = rl.NewRectangle(59, 634, 16, 16)

	dog2iimg = rl.NewRectangle(80, 657, 16, 16)
	dog2rimg = rl.NewRectangle(12, 656, 16, 16)
	dog2limg = rl.NewRectangle(59, 677, 16, 16)

	// walls & floors
	wallimg, floorimg   rl.Rectangle
	wallimgs, floorimgs []rl.Rectangle

	//stairs
	stairs1img = rl.NewRectangle(197, 166, 16, 16)
	stairs2img = rl.NewRectangle(213, 166, 16, 16)

	//chests
	chestimg = rl.NewRectangle(3, 166, 17, 17)

	//potions
	potionimgs      []rl.Rectangle
	potionemptyimgs []rl.Rectangle

	//gems jewellery coins
	gemimgs     []rl.Rectangle
	jewelryimgs []rl.Rectangle

	//weapons
	swordimgs        []rl.Rectangle
	swordimgsl       []rl.Rectangle
	daggerimgs       []rl.Rectangle
	daggerimgsl      []rl.Rectangle
	clubimgs         []rl.Rectangle
	clubimgsl        []rl.Rectangle
	scytheimgs       []rl.Rectangle
	scytheimgsl      []rl.Rectangle
	maceimgs         []rl.Rectangle
	maceimgsl        []rl.Rectangle
	wandimgs         []rl.Rectangle
	wandimgsl        []rl.Rectangle
	axeimgs          []rl.Rectangle
	axeimgsl         []rl.Rectangle
	throwingaxeimgs  []rl.Rectangle
	throwingaxeimgsl []rl.Rectangle
	spearimgs        []rl.Rectangle
	spearimgsl       []rl.Rectangle
	crossbowimgs     []rl.Rectangle
	crossbowimgsl    []rl.Rectangle
	bowimgs          []rl.Rectangle
	bowimgsl         []rl.Rectangle

	arrowimg      = rl.NewRectangle(1004, 207, 16, 16)
	quiverimg     = rl.NewRectangle(1001, 190, 16, 16)
	ninjastar1img = rl.NewRectangle(3, 186, 16, 16)
	bombimg       = rl.NewRectangle(70, 318, 16, 16)

	//legendary armor
	helmetimgs []rl.Rectangle
	bootimgs   []rl.Rectangle
	gloveimgs  []rl.Rectangle
	vestimgs   []rl.Rectangle
	robeimgs   []rl.Rectangle
	crownimgs  []rl.Rectangle

	//icons
	settingsimg = rl.NewRectangle(89, 73, 18, 18)

	//arrows
	uparrowimg    = rl.NewRectangle(19, 75, 16, 16)
	rightarrowimg = rl.NewRectangle(35, 75, 16, 16)
	downarrowimg  = rl.NewRectangle(51, 75, 16, 16)
	leftarrowimg  = rl.NewRectangle(67, 75, 16, 16)

	//boss bullets
	bullet1img = rl.NewRectangle(3, 319, 16, 16)
	bullet2img = rl.NewRectangle(19, 319, 16, 16)
	bullet3img = rl.NewRectangle(35, 319, 16, 16)
	bullet4img = rl.NewRectangle(51, 319, 16, 16)
	missileimg = rl.NewRectangle(91, 320, 16, 16)

	//objs
	campfireimg    = rl.NewRectangle(71, 166, 16, 16)
	torch1img      = rl.NewRectangle(89, 166, 16, 16)
	torch2img      = rl.NewRectangle(105, 166, 16, 16)
	torch3img      = rl.NewRectangle(121, 166, 16, 16)
	weed1img       = rl.NewRectangle(138, 167, 16, 16)
	weed2img       = rl.NewRectangle(157, 166, 16, 16)
	weed3img       = rl.NewRectangle(177, 166, 16, 16)
	weed4img       = rl.NewRectangle(347, 166, 16, 16)
	sign1img       = rl.NewRectangle(228, 166, 16, 16)
	sign2img       = rl.NewRectangle(244, 166, 16, 16)
	sign3img       = rl.NewRectangle(260, 166, 16, 16)
	bookcase1img   = rl.NewRectangle(277, 166, 16, 16)
	bookcase2img   = rl.NewRectangle(294, 166, 16, 16)
	bones1img      = rl.NewRectangle(313, 167, 16, 16)
	bones2img      = rl.NewRectangle(329, 167, 16, 16)
	grave1img      = rl.NewRectangle(365, 166, 16, 16)
	cauldronimg    = rl.NewRectangle(383, 166, 16, 16)
	key1img        = rl.NewRectangle(403, 166, 16, 16)
	key2img        = rl.NewRectangle(419, 166, 16, 16)
	key3img        = rl.NewRectangle(435, 166, 16, 16)
	food1img       = rl.NewRectangle(457, 167, 16, 16)
	food2img       = rl.NewRectangle(473, 167, 16, 16)
	food3img       = rl.NewRectangle(491, 167, 16, 16)
	food4img       = rl.NewRectangle(509, 167, 16, 16)
	food5img       = rl.NewRectangle(525, 167, 16, 16)
	food6img       = rl.NewRectangle(541, 167, 16, 16)
	food7img       = rl.NewRectangle(559, 167, 16, 16)
	spadeimg       = rl.NewRectangle(577, 168, 16, 16)
	spadeimgl      = rl.NewRectangle(743, 269, 16, 16)
	wateringcanimg = rl.NewRectangle(596, 168, 16, 16)
	bones3img      = rl.NewRectangle(613, 169, 16, 16)
	plant1img      = rl.NewRectangle(628, 169, 16, 16)
	map1img        = rl.NewRectangle(645, 169, 16, 16)
	scroll1img     = rl.NewRectangle(661, 169, 16, 16)
	scroll2img     = rl.NewRectangle(677, 169, 16, 16)
	cactusimg      = rl.NewRectangle(698, 168, 16, 16)
	grave2img      = rl.NewRectangle(722, 168, 16, 16)
	grave3img      = rl.NewRectangle(746, 168, 16, 16)
	grave4img      = rl.NewRectangle(765, 167, 18, 18)
	bones4img      = rl.NewRectangle(785, 165, 23, 23)
	bones5img      = rl.NewRectangle(813, 169, 16, 16)
	flower1img     = rl.NewRectangle(831, 170, 16, 16)
	flower2img     = rl.NewRectangle(855, 170, 16, 16)
	flower3img     = rl.NewRectangle(878, 170, 16, 16)
	weed5img       = rl.NewRectangle(898, 171, 16, 16)
	weed6img       = rl.NewRectangle(922, 171, 16, 16)
	weed7img       = rl.NewRectangle(946, 171, 16, 16)
	weed8img       = rl.NewRectangle(968, 171, 16, 16)
	weed9img       = rl.NewRectangle(991, 171, 16, 16)
	weed10img      = rl.NewRectangle(1016, 171, 16, 16)
	rocks1img      = rl.NewRectangle(1039, 172, 17, 17)
	rocks2img      = rl.NewRectangle(1063, 171, 18, 18)
	treestumpimg   = rl.NewRectangle(1086, 173, 17, 17)
	sign4img       = rl.NewRectangle(1114, 172, 16, 16)
	sign5img       = rl.NewRectangle(1140, 173, 16, 16)
	sign6img       = rl.NewRectangle(1114, 173, 16, 16)
	sign7img       = rl.NewRectangle(1167, 173, 16, 16)
	sign8img       = rl.NewRectangle(1193, 173, 16, 16)
	sign9img       = rl.NewRectangle(1217, 174, 16, 16)
	book1img       = rl.NewRectangle(1241, 173, 16, 16)
	book2img       = rl.NewRectangle(1266, 173, 16, 16)
	scroll3img     = rl.NewRectangle(1289, 172, 18, 18)
	map2img        = rl.NewRectangle(1315, 174, 18, 18)
	moneybag1img   = rl.NewRectangle(1344, 171, 16, 16)
	moneybag2img   = rl.NewRectangle(1368, 171, 16, 16)

	//backobjs
	back1img  = rl.NewRectangle(1, 245, 20, 20)
	back2img  = rl.NewRectangle(25, 245, 20, 20)
	back3img  = rl.NewRectangle(49, 241, 20, 20)
	back4img  = rl.NewRectangle(72, 241, 20, 20)
	back5img  = rl.NewRectangle(97, 240, 20, 20)
	back6img  = rl.NewRectangle(117, 242, 20, 20)
	back7img  = rl.NewRectangle(141, 242, 20, 20)
	back8img  = rl.NewRectangle(169, 241, 20, 20)
	back9img  = rl.NewRectangle(193, 241, 20, 20)
	back10img = rl.NewRectangle(220, 241, 20, 20)
	back11img = rl.NewRectangle(244, 240, 20, 20)
	back12img = rl.NewRectangle(267, 239, 20, 20)
	back13img = rl.NewRectangle(293, 242, 20, 20)
	back14img = rl.NewRectangle(319, 244, 20, 20)
	back15img = rl.NewRectangle(347, 244, 20, 20)
	back16img = rl.NewRectangle(369, 244, 20, 20)
	back17img = rl.NewRectangle(394, 245, 20, 20)
	back18mg  = rl.NewRectangle(417, 244, 20, 20)
	back19img = rl.NewRectangle(442, 244, 20, 20)
	back20img = rl.NewRectangle(466, 244, 20, 20)
	back21img = rl.NewRectangle(490, 245, 20, 20)
	back22img = rl.NewRectangle(514, 244, 20, 20)
	back23img = rl.NewRectangle(537, 244, 20, 20)
	back24img = rl.NewRectangle(562, 243, 20, 20)
	back25img = rl.NewRectangle(2, 270, 20, 20)
	back26img = rl.NewRectangle(24, 270, 20, 20)
	back27img = rl.NewRectangle(48, 266, 20, 20)
	back28img = rl.NewRectangle(72, 266, 20, 20)
	back29img = rl.NewRectangle(103, 269, 20, 20)
	back30img = rl.NewRectangle(126, 269, 20, 20)
	back31img = rl.NewRectangle(156, 270, 20, 20)
	back32img = rl.NewRectangle(186, 269, 20, 20)
	back33img = rl.NewRectangle(214, 268, 20, 20)
	back34img = rl.NewRectangle(238, 268, 20, 20)
	back35img = rl.NewRectangle(268, 264, 20, 20)
	back36img = rl.NewRectangle(293, 265, 20, 20)
	back37img = rl.NewRectangle(319, 268, 20, 20)
	back38img = rl.NewRectangle(347, 271, 20, 20)
	back39img = rl.NewRectangle(371, 272, 20, 20)
	back40img = rl.NewRectangle(395, 271, 20, 20)
	back41img = rl.NewRectangle(419, 272, 20, 20)
	back42img = rl.NewRectangle(443, 271, 20, 20)
	back43img = rl.NewRectangle(470, 270, 20, 20)
	back44img = rl.NewRectangle(497, 271, 20, 20)
	back45img = rl.NewRectangle(521, 269, 20, 20)
	back46img = rl.NewRectangle(2, 293, 20, 20)
	back47img = rl.NewRectangle(25, 293, 20, 20)
	back48img = rl.NewRectangle(50, 290, 20, 20)
	back49img = rl.NewRectangle(72, 290, 20, 20)
	back50img = rl.NewRectangle(99, 293, 20, 20)
	back51img = rl.NewRectangle(129, 294, 20, 20)
	back52img = rl.NewRectangle(157, 296, 20, 20)
	back53img = rl.NewRectangle(186, 294, 20, 20)
	back54img = rl.NewRectangle(214, 292, 20, 20)
	back55img = rl.NewRectangle(238, 292, 20, 20)
	back56img = rl.NewRectangle(267, 288, 20, 20)
	back57img = rl.NewRectangle(293, 289, 20, 20)
	back58img = rl.NewRectangle(319, 292, 20, 20)
	back59img = rl.NewRectangle(347, 295, 20, 20)
	back60img = rl.NewRectangle(371, 295, 20, 20)
	back61img = rl.NewRectangle(395, 296, 20, 20)
	back62img = rl.NewRectangle(419, 295, 20, 20)
	back63img = rl.NewRectangle(443, 294, 20, 20)

	//monsters
	monster1img  = rl.NewRectangle(1386, 220, 16, 16)
	monster2img  = rl.NewRectangle(1385, 240, 16, 16)
	monster3img  = rl.NewRectangle(1385, 261, 16, 16)
	monster4img  = rl.NewRectangle(1385, 280, 16, 16)
	monster5img  = rl.NewRectangle(1385, 300, 16, 16)
	monster6img  = rl.NewRectangle(1385, 320, 16, 16)
	monster7img  = rl.NewRectangle(1387, 340, 16, 16)
	monster8img  = rl.NewRectangle(1386, 360, 16, 16)
	monster9img  = rl.NewRectangle(1386, 380, 16, 16)
	monster10img = rl.NewRectangle(1385, 400, 16, 16)
	monster11img = rl.NewRectangle(1385, 420, 16, 16)
	monster12img = rl.NewRectangle(1385, 440, 16, 16)
	monster13img = rl.NewRectangle(1383, 460, 16, 16)
	monster15img = rl.NewRectangle(1466, 240, 16, 16)
	monster16img = rl.NewRectangle(1465, 260, 16, 16)
	monster17img = rl.NewRectangle(1463, 280, 16, 16)
	monster18img = rl.NewRectangle(1471, 300, 16, 16)
	monster19img = rl.NewRectangle(1505, 320, 16, 16)
	monster20img = rl.NewRectangle(1505, 340, 16, 16)
	monster21img = rl.NewRectangle(1486, 360, 16, 16)
	monster22img = rl.NewRectangle(1504, 380, 16, 16)
	monster23img = rl.NewRectangle(1504, 400, 16, 16)
	monster24img = rl.NewRectangle(1484, 440, 16, 16)
	monster25img = rl.NewRectangle(1527, 220, 16, 16)
	monster26img = rl.NewRectangle(1450, 488, 16, 16)
	monster27img = rl.NewRectangle(1485, 462, 16, 16)
)

// MARK: struct
type xshopitem struct {
	name, name2  string
	cost, amount int
	locked, sold bool
	img          rl.Rectangle
	color        rl.Color
	ro           float32
}
type xplayer struct {
	cnt, v1, v2, v3, v4           rl.Vector2
	rec, img, boundrec            rl.Rectangle
	color                         rl.Color
	vel, velorig, dirx, diry, ro  float32
	roomnum, hp, hpmax, dampcount int

	hpppause, burntimer, poisontimer, poisonresistancetimer, fireresistancetimer, damptimer, damppause, sicktimer, rolltimer int32

	dex, intel, str, dexorig, intelorig, strorig, luck, luckorig, fireresistance, poisonresistance, fireresistancejewel, poisonresistancejewel, fireresistancepotion, poisonresistancepotion, vampirelev, thornslev, coins, teleports int

	burning, poisoned, sick, lr, immune, speed, flametrail, rainbow, nomove, inwater bool

	ammo, weapon, object xobj
}
type xmonster struct {
	rec, img                        rl.Rectangle
	cnt                             rl.Vector2
	startx, endx, frames            float32
	move, movenum, num, hp, atktype int
	hppause, timer                  int32

	moveswitch, moveswitch2, inactiv, poisoned, burning bool

	name string
}
type xroomrec struct {
	rec, collisrec, innerrec rl.Rectangle
	cnt                      rl.Vector2
	backg                    []xbackg
	visited                  bool
}
type xroom struct {
	roomrec      []xroomrec
	num          int
	boundrec     rl.Rectangle
	xorig, yorig float32
	vis          bool
	objs         []xobj
	backobjs     []xobj
}
type xscanline struct {
	v1, v2 rl.Vector2
}
type xbackg struct {
	destrec, img rl.Rectangle
	origin       rl.Vector2
	color        rl.Color
	fade         float32
	vis          bool
}
type xobj struct {
	name, name2, name3, name4, kind string

	color, color2 rl.Color

	rec, endrec, img, imgl, origimg, boundrec, meleerangerec rl.Rectangle

	cnt, v1, v2, v3, v4 rl.Vector2

	solid, locked, inactiv, collect, nodraw, invenselect, rotates, fixed, inbound, grows, orbits, genericswitch, ident, noimg, burns, poison, legendary, questitem, msgadded, onoff, stops, line, bossbullet bool

	ro, dirx, diry, vel1, vel2, bulletsize, angle, rad, meleerange float32

	atk, scrollnum, uses, numberof, usetype, amount, armorsetnum, ability, monsternum int

	timer int32

	water []rl.Rectangle
}
type xdedmonster struct {
	v2      rl.Vector2
	timer   int32
	inactiv bool
	circles []xcircle
}
type xcircle struct {
	rad, dirx, diry, fade float32
	v2                    rl.Vector2
	color                 rl.Color
	inactiv               bool
	atk                   int
}
type xboss struct {
	rec, img, bulletimg                          rl.Rectangle
	name                                         string
	num, hp, countimg, attacktype, atk, atkspeed int
	vel, dirx, diry                              float32
	follow, switch1, onoff, inactiv              bool
	cnt                                          rl.Vector2
	followtimer                                  int32
	hppause, timer                               int32
}
type xmagic struct {
	inactiv        bool
	atk            int
	img, rec       rl.Rectangle
	dirx, diry, ro float32

	circles []xcircle
}
type xfx struct {
	img, rec rl.Rectangle
	timer    int32
	name     string
	inactiv  bool
}
type xpet struct {
	imgr, imgl, imgi, rec            rl.Rectangle
	cnt                              rl.Vector2
	name, name2                      string
	idle, offscreen, offscreenswitch bool

	frames, startxr, endxr, startxl, endxl, dirx, diry, vel float32

	timer, offscreentimer int
}

//MARK: NOTES NOTES NOTES NOTES NOTES NOTES NOTES NOTES NOTES NOTES NOTES
/*






 */
//MARK: CAMERAS CAMERAS CAMERAS CAMERAS CAMERAS CAMERAS CAMERAS CAMERAS
func cam() { //MARK: cam
	if !dev2 { //backg

		x := levboundrec.X - tilesize*4
		y := levboundrec.Y - tilesize*4

		for {
			v2 := rl.NewVector2(x+tilesize*2, y+tilesize*2)
			v22 := rl.NewVector2(x-tilesize, y-tilesize)
			rec := rl.NewRectangle(x, y, tilesize*2, tilesize*2)
			if rl.CheckCollisionPointRec(v2, borderrec) || rl.CheckCollisionPointRec(v22, borderrec) {

				rl.DrawTexturePro(imgs, wallimg, rec, origin, 0, rl.White)
			}

			x += tilesize * 2

			if x >= levboundrec.X+levboundrec.Width+tilesize*4 {
				x = levboundrec.X - tilesize*4
				y += tilesize * 2
			}

			if y >= levboundrec.Y+levboundrec.Height+tilesize*4 {
				break

			}

		}

	}
	//level layer 1 MOUSE INP
	if dev2 {
		for a := 0; a < len(level); a++ {

			rl.DrawRectangleLinesEx(level[a].boundrec, 4, rl.Blue)
			for b := 0; b < len(level[a].roomrec); b++ {
				rl.DrawRectangleLinesEx(level[a].roomrec[b].collisrec, 2, brightyellow())
				rl.DrawCircleV(level[a].roomrec[b].cnt, 4, brightred())

				txt := fmt.Sprint(level[a].num)
				rl.DrawText(txt, int32(level[a].roomrec[b].cnt.X+10), int32(level[a].roomrec[b].cnt.Y+10), 20, rl.White)
				if len(level[a].objs) > 0 {

					txt := fmt.Sprint(len(level[a].objs))
					rl.DrawText(txt, level[a].roomrec[b].collisrec.ToInt32().X+10, level[a].roomrec[b].collisrec.ToInt32().Y+10, 20, rl.White)

				}
			}
		}
	} else {

		for a := 0; a < len(visroom); a++ {

			for b := 0; b < len(visroom[a].roomrec); b++ {
				rl.DrawRectangleRec(visroom[a].roomrec[b].rec, rl.Black)

				//visited
				if rl.CheckCollisionPointRec(player.cnt, visroom[a].roomrec[b].rec) {
					visroom[a].roomrec[b].visited = true
				}

				//floor imgs
				x := visroom[a].roomrec[b].rec.X
				y := visroom[a].roomrec[b].rec.Y

				for {

					destrec := rl.NewRectangle(x, y, tilesize, tilesize)
					rl.DrawTexturePro(imgs, floorimg, destrec, origin, 0, rl.Fade(rl.DarkGray, 0.2))
					x += tilesize
					if x >= visroom[a].roomrec[b].rec.X+visroom[a].roomrec[b].rec.Width {
						x = visroom[a].roomrec[b].rec.X
						y += tilesize
					}
					if y >= visroom[a].roomrec[b].rec.Y+visroom[a].roomrec[b].rec.Height {
						break
					}
				}

				//mouse inp
				if !mapon && !settingson {
					if rl.CheckCollisionPointRec(mousev2world, visroom[a].roomrec[b].rec) && !rl.CheckCollisionPointRec(mousev2, statsrec) && !rl.CheckCollisionPointRec(mousev2, invenrec) && !rl.CheckCollisionPointRec(mousev2, msgrec) && !rl.CheckCollisionPointRec(mousev2, footerrec) {

						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
							selpoint = mousev2world
							selrec = rl.NewRectangle(selpoint.X-tilesize/8, selpoint.Y-tilesize/8, tilesize/4, tilesize/4)
							if player.speed {
								selrec = rl.NewRectangle(selpoint.X-tilesize/4, selpoint.Y-tilesize/4, tilesize/2, tilesize/2)
							}
						}

						if rl.IsMouseButtonPressed(rl.MouseRightButton) {

							if player.weapon.name != "" {
								weaponv2 = mousev2world
								useweapon()
								if soundfxon && !mute {
									rl.PlaySoundMulti(swingaud)
								}

								if weaponrotimer == 0 {
									weaponro1 = true
									weaponrotimer = weaponrotime
								}
							} else if player.object.name != "" {
								if weaponrotimer == 0 {
									weaponro1 = true
									weaponrotimer = weaponrotime
								}
								if soundfxon && !mute {
									rl.PlaySoundMulti(digaud)
								}
							} else {
								newmsg("no weapon or object equipped check inventory or find something...")
							}
							mouseclicknum++
						} else if rl.IsMouseButtonReleased(rl.MouseRightButton) {
							mouseclicknum = 0
						}

					}

				}

				if rl.CheckCollisionPointRec(mousev2world, visroom[a].roomrec[b].rec) {
					mousroomnum = b
				}

			}

		}

	}

	//level layer 2 OBJS
	for a := 0; a < len(visroom); a++ {

		//backobjs
		if len(visroom[a].backobjs) > 0 {

			for b := 0; b < len(visroom[a].backobjs); b++ {

				origin := rl.NewVector2(tilesize/2, tilesize/2)
				destrec := visroom[a].backobjs[b].rec
				destrec.X += tilesize / 2
				destrec.Y += tilesize / 2

				//obj img

				rl.DrawTexturePro(imgs, visroom[a].backobjs[b].img, destrec, origin, visroom[a].backobjs[b].ro, rl.Fade(rl.DarkGray, 0.3))

			}
		}

		//objs layer 1
		if len(visroom[a].objs) > 0 {

			for b := 0; b < len(visroom[a].objs); b++ {

				if dev2 {
					rl.DrawRectangleLinesEx(visroom[a].objs[b].rec, 4, brightorange())
					rl.DrawRectangleLinesEx(visroom[a].objs[b].boundrec, 2, brightorange())
				} else if visroom[a].objs[b].name == "water" {

					for c := 0; c < len(visroom[a].objs[b].water); c++ {

						rl.DrawRectangleRec(visroom[a].objs[b].water[c], randombluelight())
						rec2 := visroom[a].objs[b].water[c]
						rec2.X -= rFloat32(-4, 5)
						rec2.Y -= rFloat32(-4, 5)
						rl.DrawRectangleRec(rec2, rl.Fade(randombluelight(), rFloat32(0.2, 0.7)))
					}

					if !rl.CheckCollisionRecs(player.rec, visroom[a].objs[b].rec) {
						player.inwater = false
					}

				} else {
					if !visroom[a].objs[b].nodraw {

						if rl.CheckCollisionPointRec(selpoint, visroom[a].objs[b].rec) && !visroom[a].objs[b].collect {

							if selpoint.X <= visroom[a].objs[b].rec.X+visroom[a].objs[b].rec.Width/2 {
								if selpoint.Y <= visroom[a].objs[b].rec.Y+visroom[a].objs[b].rec.Height/2 {
									selpoint = rl.NewVector2(visroom[a].objs[b].rec.X, visroom[a].objs[b].rec.Y)
									selpoint.X -= tilesize / 2
									selpoint.Y -= tilesize / 2
									selrec = rl.NewRectangle(selpoint.X-tilesize/8, selpoint.Y-tilesize/8, tilesize/4, tilesize/4)
								} else {
									selpoint = rl.NewVector2(visroom[a].objs[b].rec.X, visroom[a].objs[b].rec.Y)
									selpoint.X -= tilesize / 2
									selpoint.Y += tilesize + tilesize/2
									selrec = rl.NewRectangle(selpoint.X-tilesize/8, selpoint.Y-tilesize/8, tilesize/4, tilesize/4)
								}
							} else {
								if selpoint.Y <= visroom[a].objs[b].rec.Y+visroom[a].objs[b].rec.Height/2 {
									selpoint = rl.NewVector2(visroom[a].objs[b].rec.X, visroom[a].objs[b].rec.Y)
									selpoint.X += tilesize + tilesize/2
									selpoint.Y -= tilesize / 2
									selrec = rl.NewRectangle(selpoint.X-tilesize/8, selpoint.Y-tilesize/8, tilesize/4, tilesize/4)
								} else {
									selpoint = rl.NewVector2(visroom[a].objs[b].rec.X, visroom[a].objs[b].rec.Y)
									selpoint.X += tilesize + tilesize/2
									selpoint.Y += tilesize + tilesize/2
									selrec = rl.NewRectangle(selpoint.X-tilesize/8, selpoint.Y-tilesize/8, tilesize/4, tilesize/4)
								}

							}

						}

						origin := rl.NewVector2(tilesize/2, tilesize/2)
						destrec := visroom[a].objs[b].rec
						destrec.X += tilesize / 2
						destrec.Y += tilesize / 2
						//obj img
						rl.DrawTexturePro(imgs, visroom[a].objs[b].img, destrec, origin, visroom[a].objs[b].ro, visroom[a].objs[b].color)

						if visroom[a].objs[b].name == "coin" {
							visroom[a].objs[b].img = coinimg
						}

						//obj actions
						objactions(visroom[a].num, b)

					}
				}
				if rl.CheckCollisionPointRec(mousev2world, visroom[a].objs[b].rec) && !visroom[a].objs[b].nodraw {

					//obj mouse click
					if !visroom[a].objs[b].inactiv {
						if rl.IsMouseButtonPressed(rl.MouseLeftButton) && rl.CheckCollisionRecs(visroom[a].objs[b].rec, player.boundrec) {
							clickobj(visroom[a].num, b)
						}
					}
				}
				//player obj boundrec collisions
				if rl.CheckCollisionRecs(player.rec, visroom[a].objs[b].boundrec) {
					//boundrecrec collision actions
					if !visroom[a].objs[b].inbound {
						objboundreccollisionactions(visroom[a].num, b)
					}
				} else {
					visroom[a].objs[b].msgadded = false
					visroom[a].objs[b].inbound = false
				}
				//player obj rec collisions
				if rl.CheckCollisionRecs(player.rec, visroom[a].objs[b].rec) {
					//collect objs
					if visroom[a].objs[b].collect && !visroom[a].objs[b].nodraw {
						collectobj(visroom[a].num, b)
					}
					//rec collision actions

					objcollisionactions(visroom[a].num, b)
				} else {
					if visroom[a].objs[b].name == "water" {
						visroom[a].objs[b].msgadded = false
					}
				}
			}

		}

		//objs layer 2
		if len(visroom[a].objs) > 0 {

			for b := 0; b < len(visroom[a].objs); b++ {
				if rl.CheckCollisionPointRec(mousev2world, visroom[a].objs[b].rec) && !visroom[a].objs[b].nodraw {
					//obj name txt
					txtarrow(visroom[a].objs[b].name, visroom[a].objs[b].rec)
				}
			}

		}

	}

	//fx
	if len(fx) > 0 {

		for a := 0; a < len(fx); a++ {
			if !fx[a].inactiv {

				switch fx[a].name {

				case "flame":
					rl.DrawTexturePro(imgs, fx[a].img, fx[a].rec, origin, 0, rl.Fade(randomorange(), rFloat32(0.5, 0.9)))
					rec2 := fx[a].rec
					rec2.X += rFloat32(-5, 6)
					rec2.Y += rFloat32(-5, 6)
					rl.DrawTexturePro(imgs, fx[a].img, rec2, origin, 0, rl.Fade(randomorange(), rFloat32(0.2, 0.6)))
				}

				if rl.CheckCollisionRecs(player.rec, fx[a].rec) {
					if player.fireresistance == 0 {
						player.burning = true
						player.burntimer = fps * 4
					} else {
						switch player.fireresistance {
						case 10:
							choose := rInt(1, 10)
							if choose != 1 {
								player.burning = true
								player.burntimer = fps * 4
								newmsg("OUCH! that burns >> HP -1")
							}
						case 20:
							choose := rInt(1, 10)
							if choose > 2 {
								player.burning = true
								player.burntimer = fps * 4
								newmsg("OUCH! that burns >> HP -1")
							}
						case 30:
							choose := rInt(1, 10)
							if choose > 3 {
								player.burning = true
								player.burntimer = fps * 4
								newmsg("OUCH! that burns >> HP -1")
							}
						case 40:
							choose := rInt(1, 10)
							if choose > 4 {
								player.burning = true
								player.burntimer = fps * 4
								newmsg("OUCH! that burns >> HP -1")
							}
						case 50:
							choose := rInt(1, 10)
							if choose > 5 {
								player.burning = true
								player.burntimer = fps * 4
								newmsg("OUCH! that burns >> HP -1")
							}
						case 60:
							choose := rInt(1, 10)
							if choose > 6 {
								player.burning = true
								player.burntimer = fps * 4
								newmsg("OUCH! that burns >> HP -1")
							}
						case 70:
							choose := rInt(1, 10)
							if choose > 7 {
								player.burning = true
								player.burntimer = fps * 4
								newmsg("OUCH! that burns >> HP -1")
							}
						case 80:
							choose := rInt(1, 10)
							if choose > 8 {
								player.burning = true
								player.burntimer = fps * 4
								newmsg("OUCH! that burns >> HP -1")
							}
						case 90:
							choose := rInt(1, 10)
							if choose > 9 {
								player.burning = true
								player.burntimer = fps * 4
								newmsg("OUCH! that burns >> HP -1")
							}

						}

					}

				}

				if fx[a].timer > 1 {
					fx[a].timer--
				} else if fx[a].timer == 1 {
					fx[a].inactiv = true
				}

			}
		}

	}
	//enemy bullets
	clear := false
	if len(enemybullets) > 0 && len(enemybullets) < 100 {
		for a := 0; a < len(enemybullets); a++ {
			if !enemybullets[a].inactiv {

				if enemybullets[a].noimg && !enemybullets[a].line {
					rl.DrawCircleGradient(int32(enemybullets[a].cnt.X), int32(enemybullets[a].cnt.Y), enemybullets[a].rad, enemybullets[a].color, enemybullets[a].color2)

					if !rl.CheckCollisionPointRec(enemybullets[a].cnt, borderrec) {
						enemybullets[a].inactiv = true
					}

				} else if enemybullets[a].noimg && enemybullets[a].line {

					rl.DrawLineEx(enemybullets[a].v1, enemybullets[a].v2, 12, enemybullets[a].color)

					v3 := enemybullets[a].v1
					v4 := enemybullets[a].v2

					v3.X += rFloat32(-10, 10)
					v3.Y += rFloat32(-10, 10)

					v4.X += rFloat32(-10, 10)
					v4.Y += rFloat32(-10, 10)

					rl.DrawLineEx(v3, v4, 12, rl.Fade(enemybullets[a].color, 0.4))

					enemybullets[a].v1 = boss[enemybullets[a].monsternum].cnt

					if frames%30 == 0 {
						enemybullets[a].v2.X += rFloat32(-tilesize*2, tilesize*2)
						enemybullets[a].v2.Y += rFloat32(-tilesize*2, tilesize*2)

					}

				} else {
					origin := rl.NewVector2(enemybullets[a].rec.Width/2, enemybullets[a].rec.Height/2)
					destrec := enemybullets[a].rec
					destrec.X += enemybullets[a].rec.Width / 2
					destrec.Y += enemybullets[a].rec.Height / 2

					rl.DrawTexturePro(imgs, enemybullets[a].img, destrec, origin, enemybullets[a].ro, rl.Fade(enemybullets[a].color, fadeblink))

					if !rl.CheckCollisionPointRec(enemybullets[a].cnt, borderrec) {
						enemybullets[a].inactiv = true
					}
				}

				if enemybullets[a].stops {

					if !rl.CheckCollisionPointRec(enemybullets[a].cnt, enemybullets[a].endrec) {
						enemybullets[a].cnt.X += enemybullets[a].dirx
						enemybullets[a].cnt.Y += enemybullets[a].diry
					} else {
						if !enemybullets[a].onoff {
							explode(a)
							enemybullets[a].onoff = true
						}
					}

				} else {
					enemybullets[a].cnt.X += enemybullets[a].dirx
					enemybullets[a].cnt.Y += enemybullets[a].diry
				}

				if enemybullets[a].orbits {

					enemybullets[a].angle = enemybullets[a].angle * (math.Pi / 180)

					newx := float32(math.Cos(float64(enemybullets[a].angle)))*(enemybullets[a].cnt.X-enemybullets[a].v1.X) - float32(math.Sin(float64(enemybullets[a].angle)))*(enemybullets[a].cnt.Y-enemybullets[a].v1.Y) + enemybullets[a].v1.X

					newy := float32(math.Sin(float64(enemybullets[a].angle)))*(enemybullets[a].cnt.X-enemybullets[a].v1.X) + float32(math.Cos(float64(enemybullets[a].angle)))*(enemybullets[a].cnt.Y-enemybullets[a].v1.Y) + enemybullets[a].v1.Y

					enemybullets[a].cnt = rl.NewVector2(newx, newy)

					enemybullets[a].angle += 2

					if rl.CheckCollisionRecs(enemybullets[a].rec, player.rec) {
						if player.hpppause == 0 {
							player.hp -= enemybullets[a].atk
							if soundfxon && !mute {
								rl.PlaySoundMulti(playerdamageaud)
							}
							if player.hp < 0 {
								player.hp = 0
							}
							player.hpppause = fps
						}

					}

				}

				enemybullets[a].rec = rl.NewRectangle(enemybullets[a].cnt.X-enemybullets[a].bulletsize/2, enemybullets[a].cnt.Y-enemybullets[a].bulletsize/2, enemybullets[a].bulletsize, enemybullets[a].bulletsize)

				if enemybullets[a].rotates {
					enemybullets[a].ro += 7
				}
				if enemybullets[a].grows {
					enemybullets[a].bulletsize += 0.5
				}

				if enemybullets[a].line {
					if rl.CheckCollisionPointRec(enemybullets[a].v2, player.rec) {
						if player.hpppause == 0 {
							player.hp -= enemybullets[a].atk
							if soundfxon && !mute {
								rl.PlaySoundMulti(playerdamageaud)
							}
							if player.hp < 0 {
								player.hp = 0
							}
							player.hpppause = fps
						}

					}

				} else {
					if rl.CheckCollisionRecs(player.rec, enemybullets[a].rec) {
						if enemybullets[a].burns {
							if player.fireresistance == 0 {
								player.burning = true
								player.burntimer = fps * 4
							} else {
								switch player.fireresistance {
								case 10:
									choose := rInt(1, 10)
									if choose != 1 {
										player.burning = true
										player.burntimer = fps * 4
									}
								case 20:
									choose := rInt(1, 10)
									if choose > 2 {
										player.burning = true
										player.burntimer = fps * 4
									}
								case 30:
									choose := rInt(1, 10)
									if choose > 3 {
										player.burning = true
										player.burntimer = fps * 4
									}
								case 40:
									choose := rInt(1, 10)
									if choose > 4 {
										player.burning = true
										player.burntimer = fps * 4
									}
								case 50:
									choose := rInt(1, 10)
									if choose > 5 {
										player.burning = true
										player.burntimer = fps * 4
									}
								case 60:
									choose := rInt(1, 10)
									if choose > 6 {
										player.burning = true
										player.burntimer = fps * 4
									}
								case 70:
									choose := rInt(1, 10)
									if choose > 7 {
										player.burning = true
										player.burntimer = fps * 4
									}
								case 80:
									choose := rInt(1, 10)
									if choose > 8 {
										player.burning = true
										player.burntimer = fps * 4
									}
								case 90:
									choose := rInt(1, 10)
									if choose > 9 {
										player.burning = true
										player.burntimer = fps * 4
									}

								}

							}
						}
						if enemybullets[a].poison {
							if player.poisonresistance == 0 {
								player.poisoned = true
								player.poisontimer = fps * 4
							} else {
								switch player.poisonresistance {
								case 10:
									choose := rInt(1, 10)
									if choose != 1 {
										player.poisoned = true
										player.poisontimer = fps * 4
									}
								case 20:
									choose := rInt(1, 10)
									if choose > 2 {
										player.poisoned = true
										player.poisontimer = fps * 4
									}
								case 30:
									choose := rInt(1, 10)
									if choose > 3 {
										player.poisoned = true
										player.poisontimer = fps * 4
									}
								case 40:
									choose := rInt(1, 10)
									if choose > 4 {
										player.poisoned = true
										player.poisontimer = fps * 4
									}
								case 50:
									choose := rInt(1, 10)
									if choose > 5 {
										player.poisoned = true
										player.poisontimer = fps * 4
									}
								case 60:
									choose := rInt(1, 10)
									if choose > 6 {
										player.poisoned = true
										player.poisontimer = fps * 4
									}
								case 70:
									choose := rInt(1, 10)
									if choose > 7 {
										player.poisoned = true
										player.poisontimer = fps * 4
									}
								case 80:
									choose := rInt(1, 10)
									if choose > 8 {
										player.poisoned = true
										player.poisontimer = fps * 4
									}
								case 90:
									choose := rInt(1, 10)
									if choose > 9 {
										player.poisoned = true
										player.poisontimer = fps * 4
									}

								}

							}
						}
						if player.hpppause == 0 {
							player.hp -= enemybullets[a].atk
							if soundfxon && !mute {
								rl.PlaySoundMulti(playerdamageaud)
							}
							if player.hp < 0 {
								player.hp = 0
							}
							player.hpppause = fps
							if player.thornslev > 0 {

								choose := rolldice()

								if player.thornslev == 1 {
									if choose > 4 {
										monsternum := findvismonsternum(enemybullets[a].monsternum)
										if vismonsters[monsternum].hppause == 0 {

											monsters[vismonsters[monsternum].num].hppause = fps / 2
											vismonsters[monsternum].hppause = fps / 2
											vismonsters[monsternum].hp -= enemybullets[a].atk
											monsters[vismonsters[monsternum].num].hp -= enemybullets[a].atk
											if monsters[vismonsters[monsternum].num].hp <= 0 {
												monsters[vismonsters[monsternum].num].inactiv = true
												makededmonster(monsters[vismonsters[monsternum].num].cnt)
											}
										}

										newmsg("-" + fmt.Sprint(enemybullets[a].atk) + " thorns damage...")
									}
								} else if player.thornslev == 2 {
									if choose > 2 {
										monsternum := findvismonsternum(enemybullets[a].monsternum)
										if vismonsters[monsternum].hppause == 0 {

											monsters[vismonsters[monsternum].num].hppause = fps / 2
											vismonsters[monsternum].hppause = fps / 2
											vismonsters[monsternum].hp -= enemybullets[a].atk
											monsters[vismonsters[monsternum].num].hp -= enemybullets[a].atk
											if monsters[vismonsters[monsternum].num].hp <= 0 {
												monsters[vismonsters[monsternum].num].inactiv = true
												makededmonster(monsters[vismonsters[monsternum].num].cnt)
											}
										}

										newmsg("-" + fmt.Sprint(enemybullets[a].atk) + " thorns damage...")
									}

								} else {
									monsternum := findvismonsternum(enemybullets[a].monsternum)
									if vismonsters[monsternum].hppause == 0 {

										monsters[vismonsters[monsternum].num].hppause = fps / 2
										vismonsters[monsternum].hppause = fps / 2
										vismonsters[monsternum].hp -= enemybullets[a].atk
										monsters[vismonsters[monsternum].num].hp -= enemybullets[a].atk
										if monsters[vismonsters[monsternum].num].hp <= 0 {
											monsters[vismonsters[monsternum].num].inactiv = true
											makededmonster(monsters[vismonsters[monsternum].num].cnt)
										}
									}

									newmsg("-" + fmt.Sprint(enemybullets[a].atk) + " thorns damage...")
								}

							}
						}
					}

				}

				if !rl.CheckCollisionPointRec(enemybullets[a].cnt, borderrec) && !enemybullets[a].noimg {
					enemybullets[a].inactiv = true
				}
			} else {
				clear = true
			}
		}
	} else if len(enemybullets) > 100 {
		clearenemybullets()
	}

	//monsters
	for a := 0; a < len(vismonsters); a++ {

		if rl.CheckCollisionRecs(vismonsters[a].rec, player.rec) {
			if player.hpppause == 0 {
				player.hp -= currentlevelnum
				if soundfxon && !mute {
					rl.PlaySoundMulti(playerdamageaud)
				}
				if player.hp < 0 {
					player.hp = 0
				}
				player.hpppause = fps
			}
		}

		//monster img shadow rec
		rec2 := vismonsters[a].rec
		rec2.X -= 5
		rec2.Y += 5

		rl.DrawTexturePro(imgs, vismonsters[a].img, rec2, origin, 0, rl.Black)
		if vismonsters[a].poisoned {
			rl.DrawTexturePro(imgs, vismonsters[a].img, vismonsters[a].rec, origin, 0, randomgreen())
		} else if vismonsters[a].burning {
			rl.DrawTexturePro(imgs, vismonsters[a].img, vismonsters[a].rec, origin, 0, randomorange())
		} else {
			rl.DrawTexturePro(imgs, vismonsters[a].img, vismonsters[a].rec, origin, 0, rl.White)
		}

		if rl.CheckCollisionPointRec(mousev2world, vismonsters[a].rec) {

			txtlen := rl.MeasureText(vismonsters[a].name, 20)
			rec := rl.NewRectangle(vismonsters[a].rec.X+vismonsters[a].rec.Width/2-float32(txtlen/2)-2, vismonsters[a].rec.Y-26, float32(txtlen+4), 24)

			rl.DrawRectangleRec(rec, rl.Black)

			rl.DrawText(vismonsters[a].name, vismonsters[a].rec.ToInt32().X+vismonsters[a].rec.ToInt32().Width/2-(txtlen/2), vismonsters[a].rec.ToInt32().Y-24, 20, rl.White)

		}

		if vismonsters[a].hppause != 0 || vismonsters[a].poisoned || vismonsters[a].burning {
			v2 := vismonsters[a].cnt
			v2.Y -= tilesize
			v2.X -= hpimg.Width / 2
			rl.DrawTextureRec(imgs, hpimg, v2, rl.Fade(brightred(), fadeblink))
		}
		if ghost {

			rec3 := vismonsters[a].rec
			rec3.X += rFloat32(-4, 5)
			rec3.Y += rFloat32(-4, 5)
			rl.DrawTexturePro(imgs, vismonsters[a].img, rec3, origin, 0, rl.Fade(rl.White, rFloat32(0.1, 0.5)))
		}
		if showmonshp {
			txt := fmt.Sprint(vismonsters[a].hp)
			rl.DrawText(txt, int32(vismonsters[a].cnt.X-7), int32(vismonsters[a].cnt.Y+((tilesize/3)*2)+2), txtdef, rl.Black)
			rl.DrawText(txt, int32(vismonsters[a].cnt.X-5), int32(vismonsters[a].cnt.Y+((tilesize/3)*2)), 20, brightred())
		}

		if vismonsters[a].poisoned {
			monsters[vismonsters[a].num].timer--
			if monsters[vismonsters[a].num].timer == 0 {
				monsters[vismonsters[a].num].poisoned = false
			}
			if vismonsters[a].timer%fps == 0 {
				monsters[vismonsters[a].num].hp--
				if monsters[vismonsters[a].num].hp <= 0 {
					monsters[vismonsters[a].num].inactiv = true
					makededmonster(monsters[vismonsters[a].num].cnt)
				}
			}
		}

		if vismonsters[a].burning {
			monsters[vismonsters[a].num].timer--
			if monsters[vismonsters[a].num].timer == 0 {
				monsters[vismonsters[a].num].burning = false
			}
			if vismonsters[a].timer%fps == 0 {
				monsters[vismonsters[a].num].hp--
				if monsters[vismonsters[a].num].hp <= 0 {
					monsters[vismonsters[a].num].inactiv = true
					makededmonster(monsters[vismonsters[a].num].cnt)
				}
			}
		}

	}

	//player flying weapons wands
	clear = false
	if len(activweapons) > 0 {

		for a := 0; a < len(activweapons); a++ {
			if !activweapons[a].inactiv {

				if activweapons[a].noimg {

					switch activweapons[a].usetype {

					case 4:

						for b := 1; b < len(activweapons); b++ {
							if activweapons[b].usetype == 4 {

								rl.DrawLineEx(activweapons[b].v1, activweapons[b-1].v1, 4, randombluelight())

								v3 := activweapons[b].v1
								v3.X += rFloat32(-4, 4)
								v4 := activweapons[b-1].v1
								v4.X += rFloat32(-4, 4)

								rl.DrawLineEx(v3, v4, 4, rl.Fade(randombluelight(), rFloat32(0.3, 0.8)))

							}
						}

						for b := 0; b < len(activweapons); b++ {
							if activweapons[b].usetype == 4 {
								activweapons[a].inactiv = true
								clear = true

							}
						}

					case 2:

						destrec := rl.NewRectangle(activweapons[a].v1.X, activweapons[a].v1.Y, tilesize, tilesize)

						origin2 := rl.NewVector2(tilesize/2, tilesize/2)

						rl.DrawTexturePro(imgs, frisbeeimg, destrec, origin2, activweapons[a].ro, randomcolor())

						activweapons[a].ro += 2

					case 1:
						rl.DrawLineEx(activweapons[a].v1, activweapons[a].v2, 8, randomgreen())

						activweapons[a].v1.X += activweapons[a].dirx
						activweapons[a].v2.X += activweapons[a].dirx

						activweapons[a].v1.Y += activweapons[a].diry
						activweapons[a].v2.Y += activweapons[a].diry

						if !rl.CheckCollisionPointRec(activweapons[a].v1, visiblerec) && !rl.CheckCollisionPointRec(activweapons[a].v1, visiblerec) {
							bounceweapon(a)
						}
					}
				} else {

					destrec := activweapons[a].rec
					destrec.X += destrec.Width / 2
					destrec.Y += destrec.Height / 2

					origin2 := rl.NewVector2(destrec.Width/2, destrec.Height/2)

					rl.DrawTexturePro(imgs, activweapons[a].img, destrec, origin2, activweapons[a].ro, activweapons[a].color)

					//rl.DrawRectangleLinesEx(activweapons[a].rec, 4, brightyellow())

					if activweapons[a].name == "rotates" {
						activweapons[a].ro += 10
					}

					if activweapons[a].usetype == 3 {
						if frames%8 == 0 {
							frogimg.X += 64
							activweapons[a].rec.X -= tilesize / 4
						}
						if frogimg.X > 1495 {
							frogimg.X = 1303
						}
						activweapons[a].img = frogimg
					}

					if !rl.CheckCollisionRecs(activweapons[a].rec, borderrec) {
						activweapons[a].inactiv = true
						clear = true
					}
				}

			}
		}
	}
	if clear {
		clearactivweapons()
	}
	//ded monsters
	if len(dedmonsters) > 0 {

		for a := 0; a < len(dedmonsters); a++ {
			if !dedmonsters[a].inactiv {
				for b := 0; b < len(dedmonsters[a].circles); b++ {

					rl.DrawCircle(int32(dedmonsters[a].circles[b].v2.X), int32(dedmonsters[a].circles[b].v2.Y), dedmonsters[a].circles[b].rad, rl.Fade(dedmonsters[a].circles[b].color, rFloat32(0.5, 0.9)))

					rl.DrawCircle(int32(dedmonsters[a].circles[b].v2.X+rFloat32(-5, 5)), int32(dedmonsters[a].circles[b].v2.Y+rFloat32(-5, 5)), dedmonsters[a].circles[b].rad, rl.Fade(dedmonsters[a].circles[b].color, rFloat32(0.2, 0.4)))
				}
			}
		}

	}

	//explosion
	if len(xplodecircs) > 0 {
		for a := 0; a < len(xplodecircs); a++ {
			if !xplodecircs[a].inactiv {

				rl.DrawCircle(int32(xplodecircs[a].v2.X), int32(xplodecircs[a].v2.Y), xplodecircs[a].rad, rl.Fade(xplodecircs[a].color, rFloat32(0.5, 0.9)))

				rl.DrawCircle(int32(xplodecircs[a].v2.X+rFloat32(-5, 5)), int32(xplodecircs[a].v2.Y+rFloat32(-5, 5)), xplodecircs[a].rad, rl.Fade(xplodecircs[a].color, rFloat32(0.2, 0.4)))

			}
		}

	}

	// boss
	if len(boss) > 0 {

		for a := 0; a < len(boss); a++ {

			if !boss[a].inactiv {

				rec2 := boss[a].rec
				rec2.X -= 4
				rec2.Y += 4

				if rl.CheckCollisionPointRec(mousev2world, boss[a].rec) {

					txtlen := rl.MeasureText(boss[a].name, txtdef)

					rl.DrawText(boss[a].name, int32(boss[a].cnt.X)-txtlen/2-2, boss[a].rec.ToInt32().Y-(txtdef+6), txtdef, rl.Black)
					rl.DrawText(boss[a].name, int32(boss[a].cnt.X)-txtlen/2, boss[a].rec.ToInt32().Y-(txtdef+4), txtdef, rl.White)
				}

				//boss image
				rl.DrawTexturePro(imgs, boss[a].img, rec2, origin, 0, rl.Fade(rl.Black, 0.7))

				if showmonshp {

					rl.DrawText(fmt.Sprint(boss[a].hp), int32(boss[a].cnt.X-7), boss[a].rec.ToInt32().Y+boss[a].rec.ToInt32().Height+txtdef+2, txtdef, rl.Black)

					rl.DrawText(fmt.Sprint(boss[a].hp), int32(boss[a].cnt.X-5), boss[a].rec.ToInt32().Y+boss[a].rec.ToInt32().Height+txtdef, txtdef, brightred())
				}

				switch boss[a].name {

				case "mr ghost":
					rl.DrawTexturePro(imgs, boss[a].img, boss[a].rec, origin, 0, rl.Fade(rl.White, rFloat32(0.2, 0.8)))
				default:
					rl.DrawTexturePro(imgs, boss[a].img, boss[a].rec, origin, 0, rl.White)
				}

				switch boss[a].name {
				case "mr mushroom":
					if frames%2 == 0 {
						boss[a].countimg++
					}
					if boss[a].countimg > len(mushroomboss)-1 {
						boss[a].countimg = 0
					}
					boss[a].img = mushroomboss[boss[a].countimg]
					if boss[a].dirx < 0 {
						boss[a].img = mushroombossl[boss[a].countimg]
					}
				case "mr radish":
					if frames%4 == 0 {
						boss[a].countimg++
					}
					if boss[a].countimg > len(radishboss)-1 {
						boss[a].countimg = 0
					}
					boss[a].img = radishboss[boss[a].countimg]
					if boss[a].dirx < 0 {
						boss[a].img = radishbossl[boss[a].countimg]
					}
				case "mr spike":
					if frames%4 == 0 {
						boss[a].countimg++
					}
					if boss[a].countimg > len(spikeboss)-1 {
						boss[a].countimg = 0
					}
					boss[a].img = spikeboss[boss[a].countimg]
					if boss[a].dirx < 0 {
						boss[a].img = spikebossl[boss[a].countimg]
					}
				case "mr skull":
					if frames%4 == 0 {
						boss[a].countimg++
					}
					if boss[a].countimg > len(skullboss)-1 {
						boss[a].countimg = 0
					}
					boss[a].img = skullboss[boss[a].countimg]
					if boss[a].dirx < 0 {
						boss[a].img = skullbossl[boss[a].countimg]
					}
				case "mr ghost":
					if frames%4 == 0 {
						boss[a].countimg++
					}
					if boss[a].countimg > len(ghostboss)-1 {
						boss[a].countimg = 0
					}
					boss[a].img = ghostboss[boss[a].countimg]
					if boss[a].dirx < 0 {
						boss[a].img = ghostbossl[boss[a].countimg]
					}
				case "mr reaper":
					if frames%4 == 0 {
						boss[a].countimg++
					}
					if boss[a].countimg > len(reaperboss)-1 {
						boss[a].countimg = 0
					}
					boss[a].img = reaperboss[boss[a].countimg]
					if boss[a].dirx < 0 {
						boss[a].img = reaperbossl[boss[a].countimg]
					}
				case "mr orc":
					if frames%4 == 0 {
						boss[a].countimg++
					}
					if boss[a].countimg > len(orcboss)-1 {
						boss[a].countimg = 0
					}
					boss[a].img = orcboss[boss[a].countimg]
					if boss[a].dirx < 0 {
						boss[a].img = orcbossl[boss[a].countimg]
					}
				case "mr slime":
					if frames%4 == 0 {
						boss[a].countimg++
					}
					if boss[a].countimg > len(slimeboss)-1 {
						boss[a].countimg = 0
					}
					boss[a].img = slimeboss[boss[a].countimg]
					if boss[a].dirx < 0 {
						boss[a].img = slimebossl[boss[a].countimg]
					}
				case "mr dino":
					if frames%4 == 0 {
						boss[a].countimg++
					}
					if boss[a].countimg > len(dinoboss)-1 {
						boss[a].countimg = 0
					}
					boss[a].img = dinoboss[boss[a].countimg]
					if boss[a].dirx < 0 {
						boss[a].img = dinobossl[boss[a].countimg]
					}

				}
			}
		}
	}
	if changelevelon {
		pause = true
		makelevel()
		changelevelon = false
		if gemstotal > 20 {
			makeshop()
			shopon = true
		}

	}
	//selpoint
	if selpoint != blankv2 {
		rl.DrawCircleV(selpoint, 4, brightorange())
	}

	//magic
	if magicon {
		drawmagic()
	}

	//legendary
	if uplegendaryon {
		uplegendary()
	}

	//pets
	if len(pets) > 0 {

		for a := 0; a < len(pets); a++ {
			destrec := pets[a].rec
			if pets[a].idle {

				if rolldice() == 6 {
					destrec.Y -= 2
				}
				rl.DrawTexturePro(imgs, pets[a].imgi, destrec, origin, 0, rl.White)
			} else {

				if pets[a].dirx < 0 {
					rl.DrawTexturePro(imgs, pets[a].imgl, destrec, origin, 0, rl.White)
					if frames%12 == 0 {
						pets[a].imgl.X -= pets[a].imgl.Width
						if pets[a].imgl.X < pets[a].endxl {
							pets[a].imgl.X = pets[a].startxl
						}
					}
				} else {
					rl.DrawTexturePro(imgs, pets[a].imgr, destrec, origin, 0, rl.White)
					if frames%12 == 0 {
						pets[a].imgr.X += pets[a].imgr.Width
						if pets[a].imgr.X > pets[a].endxr {
							pets[a].imgr.X = pets[a].startxr
						}
					}
				}
			}

			if rl.CheckCollisionPointRec(mousev2world, destrec) {

				txtlen := rl.MeasureText(pets[a].name, 20)
				rec := rl.NewRectangle(destrec.X+destrec.Width/2-float32(txtlen/2)-2, destrec.Y-26, float32(txtlen+4), 24)

				rl.DrawRectangleRec(rec, rl.Black)

				rl.DrawText(pets[a].name, destrec.ToInt32().X+destrec.ToInt32().Width/2-(txtlen/2), destrec.ToInt32().Y-24, 20, rl.White)

			}

			uppets()
		}

	}

	//weapon range
	if weaponrangeon {

		player.weapon.meleerangerec.X = player.cnt.X - player.weapon.meleerangerec.Width/2
		player.weapon.meleerangerec.Y = player.cnt.Y - player.weapon.meleerangerec.Height/2

		rl.DrawRectangleRec(player.weapon.meleerangerec, rl.Fade(rl.Blue, 0.2))
		rl.DrawRectangleLinesEx(player.weapon.meleerangerec, 4, rl.Fade(rl.Blue, 0.5))

		txtlen := rl.MeasureText("range", 20)
		rl.DrawText("range", int32(player.cnt.X)-txtlen/2, int32(player.weapon.meleerangerec.Y+4), 20, rl.White)

	}
	//player
	if dev2 {
		rl.DrawRectangleLinesEx(player.rec, 2, rl.Magenta)
		rl.DrawRectangleLinesEx(player.boundrec, 4, rl.Green)
		rl.DrawCircleV(player.v1, 4, brightyellow())
		rl.DrawCircleV(player.v2, 4, brightyellow())
		rl.DrawCircleV(player.v3, 4, brightyellow())
		rl.DrawCircleV(player.v4, 4, brightyellow())

	} else {
		destrec := player.rec
		destrec.X -= tilesize / 4
		destrec.Y -= tilesize / 3
		destrec.Width += tilesize / 2
		destrec.Height += tilesize / 2

		changey := rFloat32(1, 4)
		movey := false
		if !animateplayer {

			if frames%playeranimatetimer == 0 {
				destrec.Y -= changey
				movey = true
				playeranimatetimer = rInt(20, 40)
			}

			if playeremoteon {
				destrec := player.rec
				destrec.Y -= tilesize
				rl.DrawTexturePro(imgs, emoteimg, destrec, origin, 0, rl.White)
				if frames%12 == 0 {
					emoteimg.X += 32
					if emoteimg.X > emoteimgx+64 {
						emoteimg.X = emoteimgx
					}
				}

			}
		}

		if player.lr { //player right image
			rl.DrawTexturePro(imgs, playerimg, destrec, origin, player.ro, player.color)

			if player.object.name != "" { //player object right image
				destrec := player.rec
				destrec.Width = (tilesize / 4) * 3
				destrec.Height = (tilesize / 4) * 3
				origin := rl.NewVector2(destrec.Width/2, destrec.Height/2)

				switch player.object.name {
				case "spade":
					destrec.X += tilesize + tilesize/4
					destrec.Y += (tilesize / 5) * 4

					player.weapon.ro = 0
					if weaponro1 {
						player.weapon.ro = -30
						num := rInt(4, 10)
						for a := 0; a < num; a++ {
							v2 := player.cnt
							v2.X += tilesize
							v2.X += rFloat32(-tilesize/2, tilesize/2)
							v2.Y -= rFloat32(tilesize/6, tilesize)
							rl.DrawCircleV(v2, rFloat32(3, 7), rl.Black)
						}
					} else if weaponro2 {
						player.weapon.ro = -60
						num := rInt(4, 10)
						for a := 0; a < num; a++ {
							v2 := player.cnt
							v2.X += tilesize
							v2.X += rFloat32(-tilesize/2, tilesize/2)
							v2.Y -= rFloat32(tilesize/6, tilesize)
							rl.DrawCircleV(v2, rFloat32(3, 7), rl.Black)
						}
					} else if weaponro3 {
						player.weapon.ro = -90
						num := rInt(4, 10)
						for a := 0; a < num; a++ {
							v2 := player.cnt
							v2.X += tilesize
							v2.X += rFloat32(-tilesize/2, tilesize/2)
							v2.Y -= rFloat32(tilesize/6, tilesize)
							rl.DrawCircleV(v2, rFloat32(3, 7), rl.Black)
						}

					}
					player.object.img = player.object.origimg

				}

				if movey {
					destrec.Y -= changey
					movey = false
				}

				rl.DrawTexturePro(imgs, player.object.img, destrec, origin, player.weapon.ro, player.object.color)

			}

			if player.weapon.name != "" { //player right weapon
				destrec := player.rec
				origin := rl.NewVector2(destrec.Width/2, destrec.Height/2)

				switch player.weapon.name {
				case "bow":
					destrec.X += tilesize + tilesize/3
					destrec.Y += tilesize / 2
					player.weapon.ro = -45
					if weaponro1 {
						player.weapon.ro = -15
					} else if weaponro2 {
						player.weapon.ro = -30
					} else if weaponro3 {
						player.weapon.ro = -45
					}
					player.weapon.img = player.weapon.origimg
				case "crossbow":
					destrec.X += tilesize + tilesize/3
					destrec.Y += tilesize / 2
					player.weapon.ro = 0
					if weaponro1 {
						player.weapon.ro = -15
					} else if weaponro2 {
						player.weapon.ro = -30
					} else if weaponro3 {
						player.weapon.ro = -45
					}
					player.weapon.img = player.weapon.origimg
				case "sword", "dagger", "mace", "club", "scythe", "axe", "wand":
					destrec.X += tilesize + tilesize/2
					destrec.Y += tilesize / 2
					player.weapon.ro = 0
					if weaponro1 {
						player.weapon.ro = 15
					} else if weaponro2 {
						player.weapon.ro = 30
					} else if weaponro3 {
						player.weapon.ro = 45
					}
					player.weapon.img = player.weapon.origimg
				case "spear":
					destrec.X += tilesize + tilesize/3
					destrec.Y += tilesize / 2
					player.weapon.ro = 45
					if weaponro1 {
						player.weapon.ro = 15
					} else if weaponro2 {
						player.weapon.ro = 30
					} else if weaponro3 {
						player.weapon.ro = 45
					}
					player.weapon.img = player.weapon.origimg
				case "ninja star":
					destrec.Width = (tilesize / 4) * 3
					destrec.Height = (tilesize / 4) * 3

					origin = rl.NewVector2(destrec.Width/2, destrec.Height/2)
					destrec.X += tilesize + tilesize/5
					destrec.Y += (destrec.Height / 4) * 3

				case "throwing axe":
					destrec.Width = (tilesize / 4) * 3
					destrec.Height = (tilesize / 4) * 3

					origin = rl.NewVector2(destrec.Width/2, destrec.Height/2)
					destrec.X += tilesize + tilesize/5
					destrec.Y += (destrec.Height / 4) * 3

					player.weapon.img = player.weapon.origimg
				}

				if movey {
					destrec.Y -= changey
					movey = false
				}

				rl.DrawTexturePro(imgs, player.weapon.img, destrec, origin, player.weapon.ro, player.weapon.color)

				if player.weapon.rotates {
					player.weapon.ro += 5
				}

			}
		} else { //player left image
			rl.DrawTexturePro(imgs, playerlimg, destrec, origin, player.ro, player.color)

			if player.object.name != "" { //player object left image
				destrec := player.rec
				destrec.Width = (tilesize / 4) * 3
				destrec.Height = (tilesize / 4) * 3
				origin := rl.NewVector2(destrec.Width/2, destrec.Height/2)

				switch player.object.name {
				case "spade":
					destrec.X -= tilesize / 4
					destrec.Y += (tilesize / 5) * 4

					player.weapon.ro = 0
					if weaponro1 {
						player.weapon.ro = 30
						num := rInt(4, 10)
						for a := 0; a < num; a++ {
							v2 := player.cnt
							v2.X += tilesize
							v2.X += rFloat32(-tilesize/2, tilesize/2)
							v2.Y -= rFloat32(tilesize/6, tilesize)
							rl.DrawCircleV(v2, rFloat32(3, 7), rl.Black)
						}
					} else if weaponro2 {
						player.weapon.ro = 60
						num := rInt(4, 10)
						for a := 0; a < num; a++ {
							v2 := player.cnt
							v2.X += tilesize
							v2.X += rFloat32(-tilesize/2, tilesize/2)
							v2.Y -= rFloat32(tilesize/6, tilesize)
							rl.DrawCircleV(v2, rFloat32(3, 7), rl.Black)
						}
					} else if weaponro3 {
						player.weapon.ro = 90
						num := rInt(4, 10)
						for a := 0; a < num; a++ {
							v2 := player.cnt
							v2.X += tilesize
							v2.X += rFloat32(-tilesize/2, tilesize/2)
							v2.Y -= rFloat32(tilesize/6, tilesize)
							rl.DrawCircleV(v2, rFloat32(3, 7), rl.Black)
						}

					}
					player.object.img = player.object.imgl

				}

				if movey {
					destrec.Y -= changey
					movey = false
				}

				rl.DrawTexturePro(imgs, player.object.img, destrec, origin, player.weapon.ro, player.object.color)

			}

			if player.weapon.name != "" { //player left weapon
				destrec := player.rec
				origin := rl.NewVector2(destrec.Width/2, destrec.Height/2)

				switch player.weapon.name {
				case "bow":
					destrec.X -= tilesize / 3
					destrec.Y += tilesize / 2
					player.weapon.ro = 45
					if weaponro1 {
						player.weapon.ro = 15
					} else if weaponro2 {
						player.weapon.ro = 30
					} else if weaponro3 {
						player.weapon.ro = 45
					}
					player.weapon.img = player.weapon.imgl
				case "crossbow":
					destrec.X -= tilesize / 3
					destrec.Y += tilesize / 2
					player.weapon.ro = 0
					if weaponro1 {
						player.weapon.ro = 15
					} else if weaponro2 {
						player.weapon.ro = 30
					} else if weaponro3 {
						player.weapon.ro = 45
					}
					player.weapon.img = player.weapon.imgl
				case "sword", "dagger", "mace", "club", "scythe", "axe", "wand":
					destrec.X -= tilesize / 2
					destrec.Y += tilesize / 2
					player.weapon.ro = 0
					if weaponro1 {
						player.weapon.ro = -15
					} else if weaponro2 {
						player.weapon.ro = -30
					} else if weaponro3 {
						player.weapon.ro = -45
					}
					player.weapon.img = player.weapon.imgl

				case "spear":
					destrec.X -= tilesize / 3
					destrec.Y += tilesize / 2
					player.weapon.ro = -45
					if weaponro1 {
						player.weapon.ro = -15
					} else if weaponro2 {
						player.weapon.ro = -30
					} else if weaponro3 {
						player.weapon.ro = -45
					}
					player.weapon.img = player.weapon.imgl

				case "ninja star":

					destrec.Width = (tilesize / 4) * 3
					destrec.Height = (tilesize / 4) * 3

					origin = rl.NewVector2(destrec.Width/2, destrec.Height/2)

					destrec.X -= tilesize / 4
					destrec.Y += (destrec.Height / 4) * 3

				case "throwing axe":

					destrec.Width = (tilesize / 4) * 3
					destrec.Height = (tilesize / 4) * 3

					origin = rl.NewVector2(destrec.Width/2, destrec.Height/2)
					destrec.X -= tilesize / 4
					destrec.Y += (destrec.Height / 4) * 3

					player.weapon.img = player.weapon.imgl
				}
				if movey {
					destrec.Y -= changey
					movey = false
				}

				rl.DrawTexturePro(imgs, player.weapon.img, destrec, origin, player.weapon.ro, player.weapon.color)

				if player.weapon.rotates {
					player.weapon.ro -= 5
				}

			}
		}
		if player.hpppause != 0 || player.poisoned || player.burning {
			playeremoteon = false
			v2 := player.cnt
			v2.Y -= tilesize
			v2.X -= hpimg.Width / 2
			rl.DrawTextureRec(imgs, hpimg, v2, rl.Fade(brightred(), fadeblink))
		}

	}

	//borderrec
	if dev2 {
		rl.DrawRectangleLinesEx(borderrec, 4, rl.White)
		rl.DrawRectangleLinesEx(visiblerec, 8, rl.Magenta)
	}
	//lev boundrec
	if dev2 {
		rl.DrawRectangleLinesEx(levboundrec, 4, brightorange())
		rl.DrawRectangleLinesEx(backgrec, 4, rl.Green)
	}

}
func nocambackg() { //MARK: nocambackg

}
func nocam() { //MARK: nocam

	//inven
	//inven backg
	rl.DrawRectangleRec(invenrec, brightorange())
	shadowrec := invenrec
	shadowrec.X = invenrec.X + invenrec.Width
	shadowrec.Width = tilesize / 2
	rl.DrawRectangleGradientH(shadowrec.ToInt32().X, shadowrec.ToInt32().Y, shadowrec.ToInt32().Width, shadowrec.ToInt32().Height, rl.Black, rl.Blank)

	//inven boxes
	x := tilesize / 3
	y := tilesize*2 + tilesize/4

	txtlen := rl.MeasureText("backpack", txtdef)
	rl.DrawText("backpack", (invenrec.ToInt32().X+invenrec.ToInt32().Width/2)-txtlen/2, int32(y-((tilesize/2)+tilesize/10)), txtdef, rl.Black)

	jewelteleportonoff := false
	legendaryonoff := false

	for a := 0; a < len(inven); a++ {
		rec := rl.NewRectangle(x, y, tilesize, tilesize)

		if inven[a].invenselect {
			rl.DrawRectangleRec(rec, rl.Fade(brightred(), fadeblink))
		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, 0.7))
		}

		rl.DrawRectangleLinesEx(rec, 4, rl.Black)

		if inven[a].name != "" {

			if inven[a].kind == "jewel" && inven[a].usetype == 9 && inven[a].invenselect {
				jewelteleportonoff = true
			}

			if inven[a].legendary && inven[a].invenselect {
				legendaryonoff = true
			}
			rl.DrawTexturePro(imgs, inven[a].img, rec, origin, 0, inven[a].color)
			if inven[a].numberof > 0 {
				txt := fmt.Sprint(inven[a].numberof)
				rl.DrawText(txt, rec.ToInt32().X+5, rec.ToInt32().Y+5, 10, rl.White)
			}
			if inven[a].uses > 0 {

				rl.DrawRectangle(rec.ToInt32().X+rec.ToInt32().Width-txtdef, rec.ToInt32().Y+rec.ToInt32().Height-txtdef, txtdef, txtdef, rl.Black)
				txt := fmt.Sprint(inven[a].uses)
				rl.DrawText(txt, rec.ToInt32().X+rec.ToInt32().Width-(txtdef-txtdef/4), rec.ToInt32().Y+rec.ToInt32().Height-(txtdef-txtdef/8), txtdef, rl.White)

			}

			if rl.CheckCollisionPointRec(mousev2, rec) {

				//name txt
				if !destroyitemon {
					if inven[a].ident {
						if inven[a].uses > 0 {
							txthere(inven[a].name2+" - uses "+fmt.Sprint(inven[a].uses), shadowrec.X+shadowrec.Width/2, y+tilesize/3)
						} else {
							if inven[a].name2 != "" {
								txthere(inven[a].name2, shadowrec.X+shadowrec.Width/2, y+tilesize/3)
							} else {
								if inven[a].atk > 0 {
									txthere(inven[a].name+" +"+fmt.Sprint(inven[a].atk)+" damage", shadowrec.X+shadowrec.Width/2, y+tilesize/3)
								} else {
									txthere(inven[a].name, shadowrec.X+shadowrec.Width/2, y+tilesize/3)
								}
							}
						}
					} else {
						if inven[a].uses > 0 {
							txthere(inven[a].name+" - uses "+fmt.Sprint(inven[a].uses), shadowrec.X+shadowrec.Width/2, y+tilesize/3)
						} else {
							if inven[a].atk > 0 {
								txthere(inven[a].name+" +"+fmt.Sprint(inven[a].atk)+" damage", shadowrec.X+shadowrec.Width/2, y+tilesize/3)
							} else {
								txthere(inven[a].name, shadowrec.X+shadowrec.Width/2, y+tilesize/3)
							}
						}

					}

					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if inven[a].invenselect {
							inven[a].invenselect = false

							switch inven[a].kind {
							case "armor":
								switch inven[a].ability {
								case 2:
									player.vampirelev--
									newmsg("-33% chance of vampirism healing")
								case 3:
									player.thornslev--
									newmsg("-33% chance of thorns damage")
								case 4:
									player.speed = false
									player.vel = player.velorig
									newmsg("you are no longer mr speedy gonzalez")
								case 5:
									player.flametrail = false
									player.fireresistance -= 100
									newmsg("you are no longer fire resistant or leave a flaming trail")
								case 7:
									player.rainbow = false
									newmsg("back to the same humdrum color scheme...")
								}
								if armorsetactive {
									armorsetoff(a)
								}
							case "jewel":
								upjewel(false, inven[a])
							case "spade":
								player.object = xobj{}
							case "ammo":
								player.ammo = xobj{}
								activammonum = blankint
							case "weap":
								player.weapon = xobj{}
								activweaponnum = blankint
							}
						} else {

							switch inven[a].kind {
							case "armor":
								inven[a].invenselect = true
								switchinven(inven[a].name3, a)
								switch inven[a].ability {
								case 2:
									player.vampirelev++
									newmsg("+33% chance of vampirism healing")
								case 3:
									player.thornslev++
									newmsg("+33% chance of thorns damage")
								case 4:
									player.speed = true
									player.vel = player.velorig + player.velorig/2
									newmsg("you are now super speedy")
								case 5:
									player.flametrail = true
									player.fireresistance += 100
									newmsg("you leave a burning wake of destruction and are fire resistant")
								case 7:
									player.rainbow = true
									newmsg("you are now uselessly colorful...")
								}
								checkarmorset()
							case "map":
								if !questitemon {
									makelegendaryitem()
									newmsg("new active quest legendary item to find")
									inven[a] = xobj{}
									findinvennum()
								} else {
									newmsg("you are currently on a quest, complete it or abort")
								}
							case "jewel":
								inven[a].invenselect = true
								upjewel(true, inven[a])
							case "spade":
								player.weapon = xobj{}
								activweaponnum = blankint
								player.ammo = xobj{}
								activammonum = blankint
								player.object = inven[a]
								inven[a].invenselect = true
								switchinven("weap", a)
								switchinven("ammo", a)

							case "potion":
								switch inven[a].name2 {
								case "hp max+ potion":
									if player.hpmax < 99 {
										amount := rInt(5, 16)
										player.hpmax += amount
										inven[a] = xobj{}
										findinvennum()
										newmsg("maximum HP increased by " + fmt.Sprint(amount))
									} else {
										newmsg("99 is the maximum HP")
									}
								case "healing potion":
									if player.hp < player.hpmax {
										amount := rInt(5, 16)
										player.hp += amount
										if player.hp > player.hpmax {
											player.hp = player.hpmax
										}
										inven[a] = xobj{}
										findinvennum()
										newmsg("healing potion HP +" + fmt.Sprint(amount))
									} else {
										newmsg("DON'T WASTE IT >> you are at full health")
									}
								case "cure disease potion":
									if player.sick {
										player.sick = false
										player.sicktimer = 0
										player.dampcount = 0
										player.damptimer = 0
										newmsg("YAY >> you are no longer sick")
										inven[a] = xobj{}
										findinvennum()
									} else {
										newmsg("DON'T WASTE IT >> you are not currently sick")
									}
								case "resist fire potion":
									player.fireresistancepotion += inven[a].amount
									upresistances()
									inven[a].timer = potiontimer
									activpotions = append(activpotions, inven[a])

									txt := fmt.Sprint(player.fireresistance)
									if player.fireresistance > 100 {
										txt = "100"
									}
									newmsg(txt + "%" + " chance of not catching fire")

									inven[a] = xobj{}
									findinvennum()
								case "resist poison potion":

									player.poisonresistancepotion += inven[a].amount
									upresistances()
									inven[a].timer = potiontimer
									activpotions = append(activpotions, inven[a])

									txt := fmt.Sprint(player.poisonresistance)
									if player.poisonresistance > 100 {
										txt = "100"
									}
									newmsg(txt + "%" + " chance of not getting poisoned")

									inven[a] = xobj{}
									findinvennum()
								case "poison antidote potion":
									if player.poisoned {
										player.poisoned = false
										player.poisontimer = 0
										newmsg("YAY >> you are no longer posioned")
										inven[a] = xobj{}
										findinvennum()
									} else {
										newmsg("DON'T WASTE IT >> you are not currently poisoned")
									}

								}

							case "scroll":
								if !magicon {
									activscroll = inven[a]
									usescroll()
									inven[a].ident = true
									inven[a].uses--
									if inven[a].uses == 0 {
										inven[a] = xobj{}
										findinvennum()
									}

									sound2play := scrollaud[rInt(0, len(scrollaud))]
									if soundfxon && !mute {
										rl.PlaySoundMulti(sound2play)
									}
								}
							case "weap":
								player.object = xobj{}
								player.weapon = inven[a]
								inven[a].invenselect = true
								activweaponnum = a
								switchinven("weap", a)
								switchinven("spade", a)
								switch player.weapon.name {
								case "sword", "dagger", "mace", "club", "scythe", "axe":
									weaponrangeon = true
									weaponrangetimer = fps * 3

								}
							case "ammo":
								player.object = xobj{}
								player.ammo = inven[a]
								inven[a].invenselect = true
								activammonum = a
								switchinven("ammo", a)
								switchinven("spade", a)
							}

						}

					}
					if rl.IsMouseButtonPressed(rl.MouseRightButton) {

						destroyrec = rl.NewRectangle(0, rec.Y, invenrec.Width, tilesize)

						destroytimer = fps * 5
						destroynum = a
						destroyitemon = true
					}
				}
			}

		}

		y += tilesize + tilesize/4

		if a == (len(inven)/2)-1 {
			x += tilesize + tilesize/3
			y = tilesize*2 + tilesize/4
		}

	}

	if jewelteleportonoff {
		upjewelteleport()
	}

	if legendaryonoff {
		uplegendaryon = true
	} else {
		uplegendaryon = false
	}

	//destroy move item
	if destroyitemon {
		rl.DrawRectangleRec(destroyrec, randomcolor())

		switch inven[destroynum].kind {
		case "potion", "scroll", "gem", "key", "map":

			if rl.CheckCollisionPointRec(mousev2, destroyrec) {
				txthere("tick to destroy / arrow to move to belt / cross to cancel", shadowrec.X+shadowrec.Width/2, destroyrec.Y+tilesize/3)
			}

			cancelrec := rl.NewRectangle(destroyrec.X+tilesize/8, destroyrec.Y+tilesize/8, ((tilesize / 4) * 3), ((tilesize / 4) * 3))
			rl.DrawRectangleRec(cancelrec, rl.Black)

			destrec := cancelrec
			rl.DrawTexturePro(imgs, cancelimg, destrec, origin, 0, brightred())
			if rl.CheckCollisionPointRec(mousev2, cancelrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					destroyitemon = false
					destroynum = blankint
				}
			}

			tickrec := cancelrec
			tickrec.X += tilesize

			destrec = tickrec

			rl.DrawRectangleRec(tickrec, rl.Black)
			rl.DrawTexturePro(imgs, tickimg, destrec, origin, 0, rl.Green)

			if rl.CheckCollisionPointRec(mousev2, tickrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

					inven[destroynum] = xobj{}
					findinvennum()
					destroyitemon = false
					destroynum = blankint
					newmsg("item destroyed permanently however you may find something similar along the way")

				}
			}

			moverec := tickrec
			moverec.X += tilesize

			destrec = moverec

			rl.DrawRectangleRec(moverec, rl.Black)
			rl.DrawTexturePro(imgs, downarrow2img, destrec, origin, 0, rl.Blue)

			if rl.CheckCollisionPointRec(mousev2, moverec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

					if beltinvencurrentnum < len(beltinven) {
						beltinven[beltinvencurrentnum] = inven[destroynum]
						findbeltinvennum()
						inven[destroynum] = xobj{}
						findinvennum()
						destroyitemon = false
					} else {
						newmsg("belt inventory is full, right click to destroy or move items")
					}

				}
			}

		default:

			if rl.CheckCollisionPointRec(mousev2, destroyrec) {
				txthere("tick to destroy / cross to cancel", shadowrec.X+shadowrec.Width/2, destroyrec.Y+tilesize/3)
			}

			cancelrec := rl.NewRectangle(destroyrec.X+tilesize/2, destroyrec.Y+tilesize/8, ((tilesize / 4) * 3), ((tilesize / 4) * 3))
			rl.DrawRectangleRec(cancelrec, rl.Black)

			destrec := cancelrec
			rl.DrawTexturePro(imgs, cancelimg, destrec, origin, 0, brightred())
			if rl.CheckCollisionPointRec(mousev2, cancelrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					destroyitemon = false
					destroynum = blankint
				}
			}

			tickrec := cancelrec
			tickrec.X += tilesize + tilesize/3

			destrec = tickrec

			rl.DrawRectangleRec(tickrec, rl.Black)
			rl.DrawTexturePro(imgs, tickimg, destrec, origin, 0, rl.Green)

			if rl.CheckCollisionPointRec(mousev2, tickrec) {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

					if inven[destroynum].invenselect {
						switch inven[destroynum].kind {
						case "weap":
							player.weapon = xobj{}
							activweaponnum = blankint
						case "ammo":
							player.ammo = xobj{}
							activammonum = blankint
						case "armor":
							switch inven[destroynum].ability {
							case 2:
								player.vampirelev--
								newmsg("-33% chance of vampirism healing")
							case 3:
								player.thornslev--
								newmsg("-33% chance of thorns damage")
							case 4:
								player.speed = false
								player.vel = player.velorig
								newmsg("you are no longer mr speedy gonzalez")
							case 5:
								player.flametrail = false
								player.fireresistance -= 100
								newmsg("you are no longer fire resistant or leave a flaming trail")
							case 7:
								player.rainbow = false
								newmsg("back to the same humdrum color scheme...")
							}
							if armorsetactive {
								armorsetoff(destroynum)
							}
						case "spade":
							player.object = xobj{}
						}
					}

					inven[destroynum] = xobj{}
					findinvennum()
					destroyitemon = false
					destroynum = blankint
					newmsg("item destroyed permanently however you may find something similar along the way")

				}
			}

		}
	}

	// map button
	x = tilesize / 3

	rec := rl.NewRectangle(x, y, (tilesize + tilesize + tilesize/3), (tilesize/3)*2)

	txt := "map"
	txtlen = rl.MeasureText(txt, txtdef)

	if rl.CheckCollisionPointRec(mousev2, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.White, fadeblink))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			mapon = true
			pause = true
			cammap.Target = player.cnt
			cammap.Offset.X = scrwf32 / 2
			cammap.Offset.Y = scrhf32 / 2
		}

		txthere("show map", shadowrec.X+shadowrec.Width/2, y)
	}
	rl.DrawRectangleLinesEx(rec, 4, rl.Black)
	rl.DrawText(txt, rec.ToInt32().X+rec.ToInt32().Width/2-txtlen/2, int32(y+4), txtdef, rl.Black)

	// settings icon
	y += rec.Height + tilesize/4

	rec = rl.NewRectangle(x, y, tilesize, tilesize)

	if rl.CheckCollisionPointRec(mousev2, rec) {
		rl.DrawTexturePro(imgs, settingsimg, rec, origin, 0, randomcolor())
		txthere("settings", shadowrec.X+shadowrec.Width/2, y-tilesize/2)
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			settingson = true
			pause = true
		}
	} else {
		rl.DrawTexturePro(imgs, settingsimg, rec, origin, 0, rl.White)
	}

	//mute icon
	rec.X += tilesize + tilesize/4

	if rl.CheckCollisionPointRec(mousev2, rec) {
		rl.DrawTexturePro(imgs, audioicon, rec, origin, 0, randomcolor())
		txthere("mute on/off", shadowrec.X+shadowrec.Width/2, y-tilesize/2)
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if mute {
				mute = false
			} else {
				mute = true
			}

		}
	} else {
		if mute {
			rl.DrawTexturePro(imgs, audioicon, rec, origin, 0, brightred())
		} else {
			rl.DrawTexturePro(imgs, audioicon, rec, origin, 0, rl.White)
		}
	}

	//stats
	//stats backg
	rl.DrawRectangleRec(statsrec, brightorange())
	shadowrec = statsrec
	shadowrec.Width = tilesize / 2
	shadowrec.X -= shadowrec.Width

	rl.DrawRectangleGradientH(shadowrec.ToInt32().X, shadowrec.ToInt32().Y, shadowrec.ToInt32().Width, shadowrec.ToInt32().Height, rl.Blank, rl.Black)

	xl := int32(statsrec.X + tilesize/2)
	yt := int32(msgrec.Height + tilesize/8)
	rec = rl.NewRectangle(float32(xl), float32(yt), statsrec.Width, float32(txtl*2))

	//hp
	txtlen = rl.MeasureText("HP", txtl)
	rl.DrawText("HP", xl, yt, txtl, rl.Black)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		if player.hp == 1 {
			txtlen := rl.MeasureText("it's the final countdown..nanana na nanananana", txtdef)
			txthere("it's the final countdown..nanana na nanananana", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/2)))
		} else if player.hp < 4 {
			txtlen := rl.MeasureText("hit points - this would be a good time to find food or heal", txtdef)
			txthere("hit points - this would be a good time to find food or heal", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/2)))
		} else if player.hp <= player.hpmax/2 {
			txtlen := rl.MeasureText("hit points - the current outlook for your survival is poor", txtdef)
			txthere("hit points - the current outlook for your survival is poor", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/2)))
		} else if player.hp > player.hpmax/2 {
			txtlen := rl.MeasureText("hit points", txtdef)
			txthere("hit points", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/2)))
		}
	}

	yt += int32(tilesize / 2)

	txtlen = rl.MeasureText(fmt.Sprint(player.hp)+"|"+fmt.Sprint(player.hpmax), txtl2)
	txtlen += 20
	xr := int32(statsrec.X+statsrec.Width) - txtlen

	rl.DrawText(fmt.Sprint(player.hp)+"|"+fmt.Sprint(player.hpmax), xr, yt, txtl2, rl.Black)

	//str
	yt += int32(tilesize + tilesize/4)
	rec = rl.NewRectangle(float32(xl), float32(yt), statsrec.Width, float32(txtl*2))

	rl.DrawText("STR", xl, yt, txtl, rl.Black)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		txtlen := rl.MeasureText("strength - increases melee weapon damage", txtdef)
		txthere("strength - increases melee weapon damage", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/2)))

	}
	yt += int32(tilesize / 2)

	txtlen = rl.MeasureText(fmt.Sprint(player.str), txtl2)
	txtlen += 40
	xr = int32(statsrec.X+statsrec.Width) - txtlen
	rl.DrawText(fmt.Sprint(player.str), xr, yt, txtl2, rl.Black)

	//int
	yt += int32(tilesize + tilesize/4)
	rec = rl.NewRectangle(float32(xl), float32(yt), statsrec.Width, float32(txtl*2))

	rl.DrawText("INT", xl, yt, txtl, rl.Black)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		txtlen := rl.MeasureText("intelligence - increases magic damage", txtdef)
		txthere("intelligence - increases magic damage", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/2)))

	}
	yt += int32(tilesize / 2)

	txtlen = rl.MeasureText(fmt.Sprint(player.intel), txtl2)
	txtlen += 40
	xr = int32(statsrec.X+statsrec.Width) - txtlen
	rl.DrawText(fmt.Sprint(player.intel), xr, yt, txtl2, rl.Black)

	//dex
	yt += int32(tilesize + tilesize/4)
	rec = rl.NewRectangle(float32(xl), float32(yt), statsrec.Width, float32(txtl*2))
	rl.DrawText("DEX", xl, yt, txtl, rl.Black)

	if rl.CheckCollisionPointRec(mousev2, rec) {
		txtlen := rl.MeasureText("dexterity - increases ranged weapon damage", txtdef)
		txthere("dexterity - increases ranged weapon damage", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/2)))

	}

	yt += int32(tilesize / 2)

	txtlen = rl.MeasureText(fmt.Sprint(player.dex), txtl2)
	txtlen += 40
	xr = int32(statsrec.X+statsrec.Width) - txtlen
	rl.DrawText(fmt.Sprint(player.dex), xr, yt, txtl2, rl.Black)

	//luck
	yt += int32(tilesize + tilesize/4)
	rec = rl.NewRectangle(float32(xl), float32(yt), statsrec.Width, float32(txtl*2))
	rl.DrawText("LUK", xl, yt, txtl, rl.Black)

	if rl.CheckCollisionPointRec(mousev2, rec) {
		txtlen := rl.MeasureText("luck - increases probability of finding special items", txtdef)
		txthere("luck - increases probability of finding special items", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/2)))

	}

	yt += int32(tilesize / 2)

	txtlen = rl.MeasureText(fmt.Sprint(player.luck), txtl2)
	txtlen += 40
	xr = int32(statsrec.X+statsrec.Width) - txtlen
	rl.DrawText(fmt.Sprint(player.luck), xr, yt, txtl2, rl.Black)

	yt += int32(tilesize + (tilesize / 2))

	//effects icons
	xl = int32(statsrec.X + tilesize/3)
	destrec := rl.NewRectangle(float32(xl), float32(yt), tilesize, tilesize)
	if player.burning {

		rl.DrawRectangleRec(destrec, rl.Black)
		rl.DrawTexturePro(imgs, flameimg, destrec, origin, 0, randomorange())
		if rl.CheckCollisionPointRec(mousev2, destrec) {
			txtlen := rl.MeasureText("burning HP -1 per second", txtdef)
			txthere("burning HP -1 per second", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/3)))
		}
	} else {
		rl.DrawTexturePro(imgs, flameimg, destrec, origin, 0, rl.Fade(rl.DarkGray, 0.4))
		if rl.CheckCollisionPointRec(mousev2, destrec) {
			txtlen := rl.MeasureText("not burning... cool as ice", txtdef)
			txthere("not burning...cool as ice", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/3)))
		}
	}
	destrec = rl.NewRectangle(float32(xl)+tilesize+tilesize/4, float32(yt), tilesize, tilesize)
	if player.fireresistance > 0 {
		rl.DrawRectangleRec(destrec, rl.Black)
		rl.DrawTexturePro(imgs, fireresistimg, destrec, origin, 0, randomorange())
		if rl.CheckCollisionPointRec(mousev2, destrec) {
			txt := fmt.Sprint(player.fireresistance)
			if player.fireresistance > 100 {
				txt = "100"
			}
			txtlen := rl.MeasureText("fire resistance is "+txt+"%", txtdef)
			txthere("fire resistance is "+txt+"%", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/3)))
		}
	} else {
		rl.DrawTexturePro(imgs, fireresistimg, destrec, origin, 0, rl.Fade(rl.DarkGray, 0.4))
		if rl.CheckCollisionPointRec(mousev2, destrec) {
			txtlen := rl.MeasureText("you can burn baby... no fire resistance", txtdef)
			txthere("you can burn baby... no fire resistance", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/3)))
		}
	}

	yt += int32(tilesize + tilesize/4)

	destrec = rl.NewRectangle(float32(xl), float32(yt), tilesize, tilesize)
	if player.poisoned {
		rl.DrawRectangleRec(destrec, rl.Black)
		rl.DrawTexturePro(imgs, poisonedimg, destrec, origin, 0, randomgreen())
		if rl.CheckCollisionPointRec(mousev2, destrec) {
			txtlen := rl.MeasureText("poisoned HP -1 per second", txtdef)
			txthere("poisoned HP -1 per second", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/3)))

		}

	} else {
		rl.DrawTexturePro(imgs, poisonedimg, destrec, origin, 0, rl.Fade(rl.DarkGray, 0.4))
		if rl.CheckCollisionPointRec(mousev2, destrec) {
			txtlen := rl.MeasureText("not poisoned... like a mountain stream", txtdef)
			txthere("not poisoned...like a mountain stream", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/3)))
		}
	}

	destrec = rl.NewRectangle(float32(xl)+tilesize+tilesize/4, float32(yt), tilesize, tilesize)
	if player.poisonresistance > 0 {
		rl.DrawRectangleRec(destrec, rl.Black)
		rl.DrawTexturePro(imgs, poisonresistimg, destrec, origin, 0, randomgreen())
		if rl.CheckCollisionPointRec(mousev2, destrec) {
			txt := fmt.Sprint(player.poisonresistance)
			if player.poisonresistance > 100 {
				txt = "100"
			}
			txtlen := rl.MeasureText("poison resistance is "+txt+"%", txtdef)
			txthere("poison resistance is "+txt+"%", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/3)))

		}

	} else {
		rl.DrawTexturePro(imgs, poisonresistimg, destrec, origin, 0, rl.Fade(rl.DarkGray, 0.4))
		if rl.CheckCollisionPointRec(mousev2, destrec) {
			txtlen := rl.MeasureText("beware of green poisonous things... no poison resistance", txtdef)
			txthere("beware of green poisonous things... no poison resistance", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/3)))
		}
	}

	yt += int32(tilesize + tilesize/4)
	destrec = rl.NewRectangle(float32(xl), float32(yt), tilesize, tilesize)
	if player.sick {
		rl.DrawRectangleRec(destrec, rl.Black)
		rl.DrawTexturePro(imgs, sickimg, destrec, origin, 0, randomyellow())
		if rl.CheckCollisionPointRec(mousev2, destrec) {
			txtlen := rl.MeasureText("you are sick HP -1 per second", txtdef)
			txthere("you are sick HP -1 per second", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/3)))
		}
	} else {
		rl.DrawTexturePro(imgs, sickimg, destrec, origin, 0, rl.Fade(rl.DarkGray, 0.4))
		if rl.CheckCollisionPointRec(mousev2, destrec) {
			txtlen := rl.MeasureText("you are healthy... though pestilence may be near", txtdef)
			txthere("you are healthy... though pestilence may be near", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/3)))
		}
	}

	destrec = rl.NewRectangle(float32(xl)+tilesize+tilesize/4, float32(yt), tilesize, tilesize)
	if player.immune {
		rl.DrawRectangleRec(destrec, rl.Black)
		rl.DrawTexturePro(imgs, diseaseimmuneimg, destrec, origin, 0, randomyellow())
		if rl.CheckCollisionPointRec(mousev2, destrec) {
			txtlen := rl.MeasureText("step in puddles... immune to disease", txtdef)
			txthere("step in puddles... immune to disease", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/3)))
		}
	} else {
		rl.DrawTexturePro(imgs, diseaseimmuneimg, destrec, origin, 0, rl.Fade(rl.DarkGray, 0.4))
		if rl.CheckCollisionPointRec(mousev2, destrec) {
			txtlen := rl.MeasureText("don't get too cold... you are not immune to disease", txtdef)
			txthere("don't get too cold... you are not immune to disease", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/3)))
		}
	}

	//teleports
	yt += int32(tilesize + tilesize/2)
	teleportsrec := rl.NewRectangle(statsrec.X, float32(yt), statsrec.Width, tilesize)
	if rl.CheckCollisionPointRec(mousev2, teleportsrec) {
		if teleporton {
			txtlen := rl.MeasureText("teleports used on map to fast travel to visited areas", txtdef)
			txthere("teleports used on map to fast travel to visited areas", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/3)))
		} else {
			txtlen := rl.MeasureText("teleports disabled turn on in settings", txtdef)
			txthere("teleports disabled turn on in settings", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/3)))
		}
	}

	xteleport := statsrec.X + (statsrec.Width / 2) - (tilesize + tilesize/3)
	destrec = rl.NewRectangle(xteleport, float32(yt), tilesize, tilesize)
	origin2 := rl.NewVector2(tilesize/2, tilesize/2)
	destrec.X += tilesize / 2
	destrec.Y += tilesize / 2
	if teleporton {
		if rolldice() == 6 {
			rl.DrawTexturePro(imgs, teleportimg, destrec, origin2, teleportimgro, randomcolor())
		} else {
			rl.DrawTexturePro(imgs, teleportimg, destrec, origin2, teleportimgro, rl.White)
		}
	} else {
		rl.DrawTexturePro(imgs, teleportimg, destrec, origin2, teleportimgro, brightred())
	}
	teleportimgro++
	txt = "x" + fmt.Sprint(player.teleports)
	rl.DrawText(txt, int32(destrec.X-tilesize/2+destrec.Width+tilesize/8), int32(destrec.Y-tilesize/2+tilesize/8), txtl2, rl.Black)

	//coins
	yt += int32(tilesize + tilesize/2)

	coinrec := rl.NewRectangle(statsrec.X, float32(yt), statsrec.Width, tilesize)
	if rl.CheckCollisionPointRec(mousev2, coinrec) {
		txtlen := rl.MeasureText("coins used to replenish health and increase stats", txtdef)
		txthere("coins used to replenish health and increase stats", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yt+int32(tilesize/3)))
	}
	xcoin := statsrec.X + (statsrec.Width / 2) - (tilesize + tilesize/3)
	destrec = rl.NewRectangle(xcoin, float32(yt), tilesize, tilesize)
	rl.DrawTexturePro(imgs, coinimg, destrec, origin, 0, rl.White)

	txt = "x" + fmt.Sprint(player.coins)
	rl.DrawText(txt, int32(destrec.X+destrec.Width+tilesize/8), int32(destrec.Y+tilesize/8), txtl2, rl.Black)

	//map quest arrow
	yarrow := destrec.Y + tilesize*4
	x = statsrec.X + (statsrec.Width - (statsrec.Width / 8))

	if questitemon {

		ytxt := int32(yarrow - ((tilesize / 2) + (tilesize * 2)))
		xtxt := int32(statsrec.X + statsrec.Width/2)
		txtlen = rl.MeasureText("quest item", txtdef)
		rl.DrawText("quest item", xtxt-txtlen/2, ytxt, txtdef, rl.Black)

		ro := angle2points(player.cnt, questitemv2)
		destrec := rl.NewRectangle(x, float32(yarrow), tilesize*2, tilesize*2)
		origin := rl.NewVector2(destrec.Width/2, destrec.Height/2)
		destrec.X -= destrec.Width / 2
		destrec.Y -= destrec.Height / 2

		rl.DrawTexturePro(imgs, uparrowimg, destrec, origin, ro, rl.White)

		//	rl.DrawRectangleLinesEx(destrec, 2, brightred())

		x = statsrec.X + tilesize/3

		rec = rl.NewRectangle(x, yarrow+tilesize/4, (tilesize + tilesize + tilesize/3), (tilesize/3)*2)

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(rl.White, fadeblink))
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

				for a := 0; a < len(level[questitemroomnum].objs); a++ {
					if level[questitemroomnum].objs[a].questitem {
						level[questitemroomnum].objs = remobj(level[questitemroomnum].objs, a)
						break
					}
				}
				questitemon = false
				questitemv2 = blankv2
				newmsg("you have ended that quest, one less legendary piece of equipment to find... sigh")
			}

			txtlen = rl.MeasureText("end quest", 20)
			txthere("end quest", statsrec.X-float32(txtlen)-tilesize-tilesize/2, float32(yarrow+tilesize/4))

		}
		txt := "abort"
		txtlen = rl.MeasureText(txt, txtdef)
		rl.DrawRectangleLinesEx(rec, 4, rl.Black)
		rl.DrawText(txt, rec.ToInt32().X+rec.ToInt32().Width/2-txtlen/2, int32(yarrow+tilesize/4+4), txtdef, rl.Black)

	}

	//msg
	rl.DrawRectangleRec(msgrec, brightorange())
	shadowrec = msgrec
	shadowrec.Y += msgrec.Height
	shadowrec.Height = tilesize / 2
	rl.DrawRectangleGradientV(shadowrec.ToInt32().X, shadowrec.ToInt32().Y, shadowrec.ToInt32().Width, shadowrec.ToInt32().Height, rl.Black, rl.Blank)

	//close game
	closerec := rl.NewRectangle(scrwf32-(tilesize/2+tilesize/8), tilesize/8, tilesize/2, tilesize/2)
	rl.DrawRectangleRec(closerec, rl.Black)
	rl.DrawLine(closerec.ToInt32().X, closerec.ToInt32().Y, closerec.ToInt32().X+int32(tilesize/2), closerec.ToInt32().Y+int32(tilesize/2), brightorange())
	rl.DrawLine(closerec.ToInt32().X+int32(tilesize/2), closerec.ToInt32().Y, closerec.ToInt32().X, closerec.ToInt32().Y+int32(tilesize/2), brightorange())
	if rl.CheckCollisionPointRec(mousev2, closerec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			pause = true
			endgamewindow = true
		}
	}

	if newmsgtimer > 0 {
		destrec := rl.NewRectangle(msgrec.X-tilesize, msgrec.Y+5, tilesize, tilesize)
		rl.DrawTexturePro(imgs, newmsgimg, destrec, origin, 0, randomred())
	}

	if len(msgs) > 0 && newmsgtimer > 0 {
		rl.DrawText(msgs[len(msgs)-1], msgrec.ToInt32().X+10, msgrec.ToInt32().Y+10, txtl, rl.Black)
	}

	//footer
	rl.DrawRectangleRec(footerrec, brightorange())
	shadowrec = footerrec
	shadowrec.Y -= tilesize / 2
	shadowrec.Height = tilesize / 2
	rl.DrawRectangleGradientV(shadowrec.ToInt32().X, shadowrec.ToInt32().Y, shadowrec.ToInt32().Width, shadowrec.ToInt32().Height, rl.Blank, rl.Black)

	//RUNTIMER
	footy := scrh - (txtdef + ((txtdef / 3) * 2))
	txtlen = rl.MeasureText("run time : ", txtdef)
	rl.DrawText("run time : ", txtdef, footy, txtdef, rl.Black)

	runmintxt := fmt.Sprint(runmin)

	if runmin == 0 {
		runmintxt = "00"
	} else if runmin > 0 && runmin < 10 {
		runmintxt = "0" + fmt.Sprint(runmin)
	}

	runsectxt := fmt.Sprint(runsecs)

	if runsecs == 0 {
		runsectxt = "00"
	} else if runsecs > 0 && runsecs < 10 {
		runsectxt = "0" + fmt.Sprint(runsecs)
	}

	runtimetxt := runmintxt + ":" + runsectxt

	rl.DrawText(runtimetxt, txtlen+txtdef, footy, txtdef, rl.Black)

	txtlen = rl.MeasureText("run time : 00:00  ", txtdef)
	txtlen += txtdef

	//boss kills
	rl.DrawText("kills till boss : ", txtlen, footy, txtdef, rl.Black)
	txtlen += rl.MeasureText("kills till boss : ", txtdef)
	rl.DrawText(fmt.Sprint(nextbossnum-monsterkills), txtlen, footy, txtdef, rl.Black)

	//MARK: BELT BELT BELT BELT BELT BELT BELT BELT BELT BELT BELT BELT
	//quick items
	quickx := scrwf32 / 2

	length := float32(len(beltinven) / 2)

	quickx -= tilesize*length + (4 * length)
	quickrec := rl.NewRectangle(quickx, footerrec.Y+1, tilesize, tilesize)

	txtx := float32(0)
	for a := 0; a < len(beltinven); a++ {
		rl.DrawRectangleRec(quickrec, rl.Fade(rl.Black, 0.7))
		rl.DrawRectangleLinesEx(quickrec, 4, rl.Black)

		if beltinven[a].name != "" {
			rl.DrawTexturePro(imgs, beltinven[a].img, quickrec, origin, 0, beltinven[a].color)
			if beltinven[a].uses > 0 {
				rl.DrawRectangle(quickrec.ToInt32().X+quickrec.ToInt32().Width-txtdef, quickrec.ToInt32().Y+quickrec.ToInt32().Height-txtdef, txtdef, txtdef, rl.Black)
				txt := fmt.Sprint(beltinven[a].uses)
				rl.DrawText(txt, quickrec.ToInt32().X+quickrec.ToInt32().Width-(txtdef-txtdef/4), quickrec.ToInt32().Y+quickrec.ToInt32().Height-(txtdef-txtdef/8), txtdef, rl.White)
			}

			if rl.CheckCollisionPointRec(mousev2, quickrec) {

				//name txt
				if !destroybeltitemon {
					if beltinven[a].ident {
						if beltinven[a].uses > 0 {
							txthere(beltinven[a].name2+" - uses "+fmt.Sprint(beltinven[a].uses), quickrec.X, shadowrec.Y-20)
						} else {
							if beltinven[a].name2 != "" {
								txthere(beltinven[a].name2, quickrec.X, shadowrec.Y-20)
							} else {
								txthere(beltinven[a].name, quickrec.X, shadowrec.Y-20)
							}
						}
					} else {
						if beltinven[a].uses > 0 {
							txthere(beltinven[a].name+" - uses "+fmt.Sprint(beltinven[a].uses), quickrec.X, shadowrec.Y-20)
						} else {
							txthere(beltinven[a].name, quickrec.X, shadowrec.Y-20)
						}

					}
				}

				if rl.IsMouseButtonPressed(rl.MouseLeftButton) && !destroybeltitemon {
					switch beltinven[a].kind {

					case "map":
						if !questitemon {
							makelegendaryitem()
							newmsg("new active quest legendary item to find")
							inven[a] = xobj{}
							findinvennum()
						} else {
							newmsg("you are currently on a quest, complete it or abort")
						}
					case "potion":
						switch beltinven[a].name2 {
						case "hp max+ potion":
							if player.hpmax < 99 {
								player.hpmax++
								beltinven[a] = xobj{}
								findbeltinvennum()
								newmsg("maximum HP increased by 1")
							} else {
								newmsg("99 is the maximum HP")
							}
						case "healing potion":
							if player.hp < player.hpmax {
								player.hp++
								beltinven[a] = xobj{}
								findbeltinvennum()
								newmsg("healing potion HP+1")
							} else {
								newmsg("DON'T WASTE IT >> you are at full health")
							}
						case "cure disease potion":
							if player.sick {
								player.sick = false
								player.sicktimer = 0
								player.dampcount = 0
								player.damptimer = 0
								newmsg("YAY >> you are no longer sick")
								beltinven[a] = xobj{}
								findbeltinvennum()
							} else {
								newmsg("DON'T WASTE IT >> you are not currently sick")
							}
						case "resist fire potion":
							player.fireresistancepotion += beltinven[a].amount
							upresistances()
							beltinven[a].timer = potiontimer
							activpotions = append(activpotions, beltinven[a])

							txt := fmt.Sprint(player.fireresistance)
							if player.fireresistance > 100 {
								txt = "100"
							}
							newmsg(txt + "%" + " chance of not catching fire")

							beltinven[a] = xobj{}
							findbeltinvennum()
						case "resist poison potion":

							player.poisonresistancepotion += beltinven[a].amount
							upresistances()
							beltinven[a].timer = potiontimer
							activpotions = append(activpotions, beltinven[a])

							txt := fmt.Sprint(player.poisonresistance)
							if player.poisonresistance > 100 {
								txt = "100"
							}
							newmsg(txt + "%" + " chance of not getting poisoned")

							beltinven[a] = xobj{}
							findbeltinvennum()
						case "poison antidote potion":
							if player.poisoned {
								player.poisoned = false
								player.poisontimer = 0
								newmsg("YAY >> you are no longer posioned")
								beltinven[a] = xobj{}
								findbeltinvennum()
							} else {
								newmsg("DON'T WASTE IT >> you are not currently poisoned")
							}

						}

					case "scroll":
						if !magicon {
							activscroll = beltinven[a]
							usescroll()
							beltinven[a].ident = true
							beltinven[a].uses--
							if beltinven[a].uses == 0 {
								beltinven[a] = xobj{}
								findbeltinvennum()
							}

							sound2play := scrollaud[rInt(0, len(scrollaud))]
							if soundfxon && !mute {
								rl.PlaySoundMulti(sound2play)
							}
						}

					}

				}
				if rl.IsMouseButtonPressed(rl.MouseRightButton) && !destroybeltitemon {
					destroybeltrec = rl.NewRectangle(quickrec.X, quickrec.Y, invenrec.Width, tilesize)

					destroybelttimer = fps * 5
					destroybeltnum = a
					destroybeltitemon = true
				}

			}

		}

		quickrec.X += tilesize + 4
		txtx = quickrec.X
	}

	//destroy belt item
	if destroybeltitemon {
		rl.DrawRectangleRec(destroybeltrec, randomcolor())

		if rl.CheckCollisionPointRec(mousev2, destroybeltrec) {

			txthere("tick to destroy / arrow to move to belt / cross to cancel", quickrec.X, shadowrec.Y-20)

		}

		cancelrec := rl.NewRectangle(destroybeltrec.X+tilesize/8, destroybeltrec.Y+tilesize/8, ((tilesize / 4) * 3), ((tilesize / 4) * 3))
		rl.DrawRectangleRec(cancelrec, rl.Black)

		destrec := cancelrec
		rl.DrawTexturePro(imgs, cancelimg, destrec, origin, 0, brightred())
		if rl.CheckCollisionPointRec(mousev2, cancelrec) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				destroybeltitemon = false
				destroybeltnum = blankint
			}
		}

		tickrec := cancelrec
		tickrec.X += tilesize

		destrec = tickrec

		rl.DrawRectangleRec(tickrec, rl.Black)
		rl.DrawTexturePro(imgs, tickimg, destrec, origin, 0, rl.Green)

		if rl.CheckCollisionPointRec(mousev2, tickrec) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

				beltinven[destroybeltnum] = xobj{}
				findbeltinvennum()
				destroybeltitemon = false
				destroybeltnum = blankint
				newmsg("item destroyed permanently however you may find something similar along the way")

			}
		}

		moverec := tickrec
		moverec.X += tilesize

		destrec = moverec

		rl.DrawRectangleRec(moverec, rl.Black)
		rl.DrawTexturePro(imgs, uparrow2img, destrec, origin, 0, rl.Blue)

		if rl.CheckCollisionPointRec(mousev2, moverec) {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

				if invencurrentnum < len(inven) {
					inven[invencurrentnum] = beltinven[destroybeltnum]
					findinvennum()
					beltinven[destroybeltnum] = xobj{}
					findbeltinvennum()
					destroybeltitemon = false
				} else {
					newmsg("backpack inventory is full, right click to destroy or move items")
				}

			}
		}

	}

	txtx += tilesize / 4

	rl.DrawText("belt", int32(txtx), quickrec.ToInt32().Y+int32(tilesize)/3, txtdef, rl.Black)

	// stuck button
	x = txtx + tilesize + tilesize/2

	rec = rl.NewRectangle(x, quickrec.Y+tilesize/8, (tilesize + tilesize + tilesize/3), (tilesize/3)*2)

	txt = "stuck..."
	txtlen = rl.MeasureText(txt, txtdef)

	if rl.CheckCollisionPointRec(mousev2, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.White, fadeblink))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			changelevelon = true
			score -= 200
		}

		txtlen2 := rl.MeasureText("ends the level and moves to next", txtdef)

		txthere("ends the level and moves to next", rec.X-float32(txtlen2), y)
	}
	rl.DrawRectangleLinesEx(rec, 4, rl.Black)
	rl.DrawText(txt, rec.ToInt32().X+rec.ToInt32().Width/2-txtlen/2, rec.ToInt32().Y+4, txtdef, rl.Black)

	// help button
	x = rec.X + rec.Width + tilesize/4

	rec = rl.NewRectangle(x, quickrec.Y+tilesize/8, (tilesize + tilesize + tilesize/3), (tilesize/3)*2)

	txt = "help"
	txtlen = rl.MeasureText(txt, txtdef)

	if rl.CheckCollisionPointRec(mousev2, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.White, fadeblink))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			pause = true
			helpon = true
		}

		txtlen2 := rl.MeasureText("a little help", txtdef)

		txthere("a little help", rec.X-float32(txtlen2), y)
	}
	rl.DrawRectangleLinesEx(rec, 4, rl.Black)
	rl.DrawText(txt, rec.ToInt32().X+rec.ToInt32().Width/2-txtlen/2, rec.ToInt32().Y+4, txtdef, rl.Black)

	//cleared monsters
	if clearedmonsters {
		drawclearedlevel()
	}

	//blood
	if len(blood) > 0 {
		drawblood()
	}

	//map settings background
	if mapon || settingson || displaymsgs || shopon || scoreon || helpon {
		//map backg
		rl.DrawRectangle(0, 0, scrw, scrh, brightorange())
	}

	//intro background
	if introon {

		rl.DrawRectangle(0, 0, scrw, scrh, rl.Black)
	}

	if displaymsgs {
		drawmsgs()
	}
	//scanlines
	if scan && !dev2 {
		for a := 0; a < len(scanlines); a++ {
			rl.DrawLineV(scanlines[a].v1, scanlines[a].v2, rl.Fade(rl.Black, 0.7))
			scanlines[a].v1.Y += 1
			scanlines[a].v2.Y = scanlines[a].v1.Y

			if scanlines[a].v1.Y == scrhf32 {
				scanlines[a].v1.Y = 0
				scanlines[a].v2.Y = 0
			}
		}
	}

	/*
		//closewin
		if closewinloc(scrwf32-tilesize*2, tilesize, brightred(), brightyellow()) {
			rl.CloseWindow()
		}
	*/

	if teston {
		drawtest()
	}

}
func nocamMap() { //MARK: nocamMap

	//map
	if mapon {

		rl.DrawText("map", 20, 10, 80, rl.Black)

		x3 := float32(40)
		y3 := float32(150)

		v2 := rl.NewVector2(x3, y3)

		rl.DrawCircleV(v2, 15, rl.Fade(rl.Green, fadeblink))
		rl.DrawText("player", int32(v2.X+30), int32(v2.Y-15), 30, rl.Black)

		v2.Y += 50
		rl.DrawCircleV(v2, 15, rl.Fade(brightred(), fadeblink))
		rl.DrawText("staircase", int32(v2.X+30), int32(v2.Y-15), 30, rl.Black)

		v2.Y += 50
		rl.DrawCircleV(v2, 15, rl.Fade(rl.White, fadeblink))
		rl.DrawText("teleport here", int32(v2.X+30), int32(v2.Y-15), 30, rl.Black)

		v2.Y += 50
		visitedrec := rl.NewRectangle(v2.X-tilesize/4, v2.Y-tilesize/4, tilesize/2, tilesize/2)
		rl.DrawRectangleRec(visitedrec, rl.Black)
		rl.DrawRectangleRec(visitedrec, rl.Fade(rl.DarkBlue, 0.3))
		rl.DrawText("visited area", int32(v2.X+30), int32(v2.Y-15), 30, rl.Black)

		rec := rl.NewRectangle(scrwf32-(tilesize*3), (tilesize*2)+tilesize/2, tilesize, tilesize)

		txtlen := rl.MeasureText("move map", txtdef)
		txtx := rec.ToInt32().X + int32(tilesize/2)
		rl.DrawText("move map", txtx-txtlen/2, rec.ToInt32().Y-int32((tilesize*2)-(tilesize/3)), txtdef, rl.Black)
		txty := rec.Y

		rec.Y -= tilesize
		rl.DrawTexturePro(imgs, uparrowimg, rec, origin, 0, rl.Black)
		if rl.CheckCollisionPointRec(mousev2, rec) {
			if rl.IsMouseButtonDown(rl.MouseLeftButton) {
				cammap.Offset.Y += 10
			}
			rl.DrawTexturePro(imgs, uparrowimg, rec, origin, 0, rl.Fade(brightred(), fadeblink))
		} else {
			rl.DrawTexturePro(imgs, uparrowimg, rec, origin, 0, rl.White)
		}

		rec.Y += tilesize * 2
		rl.DrawTexturePro(imgs, downarrowimg, rec, origin, 0, rl.Black)
		if rl.CheckCollisionPointRec(mousev2, rec) {
			if rl.IsMouseButtonDown(rl.MouseLeftButton) {
				cammap.Offset.Y -= 10
			}
			rl.DrawTexturePro(imgs, downarrowimg, rec, origin, 0, rl.Fade(brightred(), fadeblink))
		} else {
			rl.DrawTexturePro(imgs, downarrowimg, rec, origin, 0, rl.White)
		}

		rec.Y -= tilesize
		rec.X -= tilesize
		rl.DrawTexturePro(imgs, leftarrowimg, rec, origin, 0, rl.Black)
		if rl.CheckCollisionPointRec(mousev2, rec) {
			if rl.IsMouseButtonDown(rl.MouseLeftButton) {
				cammap.Offset.X -= 10
			}
			rl.DrawTexturePro(imgs, leftarrowimg, rec, origin, 0, rl.Fade(brightred(), fadeblink))
		} else {
			rl.DrawTexturePro(imgs, leftarrowimg, rec, origin, 0, rl.White)
		}

		rec.X += tilesize * 2
		rl.DrawTexturePro(imgs, rightarrowimg, rec, origin, 0, rl.Black)
		if rl.CheckCollisionPointRec(mousev2, rec) {
			if rl.IsMouseButtonDown(rl.MouseLeftButton) {
				cammap.Offset.X += 10
			}
			rl.DrawTexturePro(imgs, rightarrowimg, rec, origin, 0, rl.Fade(brightred(), fadeblink))
		} else {
			rl.DrawTexturePro(imgs, rightarrowimg, rec, origin, 0, rl.White)
		}

		txtlen = rl.MeasureText("zoom", txtdef)

		rl.DrawText("zoom", txtx-txtlen/3, int32(txty)+int32((tilesize*3)+tilesize/10), txtdef, rl.Black)

		rec.X = float32(txtx - txtlen/2 - int32((tilesize/4)*3))
		rec.Y = txty + tilesize*3

		rec.Width = tilesize / 2
		rec.Height = tilesize / 2

		x := rec.X

		rl.DrawRectangleRec(rec, rl.Black)
		if rl.CheckCollisionPointRec(mousev2, rec) {
			if rl.IsMouseButtonDown(rl.MouseLeftButton) {
				if cammap.Zoom > 0.05 {
					cammap.Zoom -= 0.01
				}
			}
		}

		rl.DrawText("-", rec.ToInt32().X+int32(tilesize/5), rec.ToInt32().Y+int32(tilesize/9), txtdef, rl.White)

		rec.X += tilesize * 2

		rl.DrawRectangleRec(rec, rl.Black)
		if rl.CheckCollisionPointRec(mousev2, rec) {
			if rl.IsMouseButtonDown(rl.MouseLeftButton) {
				if cammap.Zoom < 0.5 {
					cammap.Zoom += 0.01
				}
			}
		}

		rl.DrawText("+", rec.ToInt32().X+int32(tilesize/5), rec.ToInt32().Y+int32(tilesize/9), txtdef, rl.White)

		y := rec.Y + tilesize*2
		x += tilesize / 4

		rec = rl.NewRectangle(x, y, (tilesize + tilesize + tilesize/3), (tilesize/3)*2)

		txt := "close"
		txtlen = rl.MeasureText(txt, txtdef)

		if rl.CheckCollisionPointRec(mousev2, rec) {
			rl.DrawRectangleRec(rec, rl.Fade(rl.White, fadeblink))
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				mapon = false
				pause = false
			}
		}
		rl.DrawRectangleLinesEx(rec, 4, rl.Black)
		rl.DrawText(txt, rec.ToInt32().X+rec.ToInt32().Width/2-txtlen/2, int32(y)+4, txtdef, rl.Black)

	}

}

func devui() { //MARK: devui

	siderec := rl.NewRectangle(0, 0, 300, scrhf32)

	rl.DrawRectangleRec(siderec, rl.Fade(rl.Green, 0.5))

	x := int32(siderec.X + 10)
	y := int32(10)

	txt := "mousev2.X"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(mousev2.X)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "mousev2.Y"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(mousev2.Y)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "scrwint"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(scrwint)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "scrhint"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(scrhint)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "len level"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(len(level))
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "player roomnum"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(player.roomnum)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "len activeweapons"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(len(activweapons))
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "len enemybullets"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(len(enemybullets))
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "magicon"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(magicon)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "player emote timer"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(playeremotetimer)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "player damp count"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(player.dampcount)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "len monsters"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(len(monsters))
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "player obj name"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = player.object.name
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "player poisonresistance"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(player.poisonresistance)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "player vampirelev"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(player.vampirelev)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "player thornslev"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(player.thornslev)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "len msgs"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(len(msgs))
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "camera zoom"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(camera.Zoom)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "current levelnum"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(currentlevelnum)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "gemstotal"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(gemstotal)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "scoreontimer"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(scoreontimer)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "player vel"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(player.vel)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

	txt = "monsternumlevel"
	rl.DrawText(txt, x, y, txts, rl.White)
	x += rl.MeasureText(txt, txts) + txts
	txt = fmt.Sprint(monsternumlevel)
	rl.DrawText(txt, x, y, txts, rl.White)
	x = int32(siderec.X + 10)
	y += txts

}

//MARK: UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE UPDATE

func update() { //MARK: update

	inp()
	timers2()

	if !pause {
		timers()
		uprooms()

		upplayer()
		if len(activweapons) > 0 {
			upweapons()
		}
		if len(dedmonsters) > 0 {
			updedmonsters()
		}
		if len(xplodecircs) > 0 {
			upexplosions()
		}
		if len(boss) > 0 {
			upboss()
		}
		if len(activpotions) > 0 {
			uppotions()
		}

		upmonsters()
		animate()
		upcams()

	}

}
func upboss() { //MARK: upboss

	for a := 0; a < len(boss); a++ {
		if !boss[a].inactiv {
			if boss[a].hppause > 0 {
				boss[a].hppause--
			}

			if boss[a].followtimer > 0 {
				boss[a].followtimer--
				if boss[a].followtimer == 0 {
					boss[a].follow = true
				}
			}

			if boss[a].follow {

				xdiff := absdiff32(boss[a].cnt.X, player.cnt.X)
				ydiff := absdiff32(boss[a].cnt.Y, player.cnt.Y)

				if xdiff > ydiff {
					boss[a].dirx = boss[a].vel
					if player.cnt.X < boss[a].cnt.X {
						boss[a].dirx = -boss[a].dirx
					}
					ychange := xdiff / boss[a].vel
					boss[a].diry = ydiff / ychange
					if player.cnt.Y < boss[a].cnt.Y {
						boss[a].diry = -boss[a].diry
					}
				} else {
					boss[a].diry = boss[a].vel
					if player.cnt.Y < boss[a].cnt.Y {
						boss[a].diry = -boss[a].diry
					}
					xchange := ydiff / boss[a].vel
					boss[a].dirx = xdiff / xchange
					if player.cnt.X < boss[a].cnt.X {
						boss[a].dirx = -boss[a].dirx
					}

				}

				checkv1x := boss[a].cnt
				checkv1x.X += boss[a].dirx

				checkv1x.X -= tilesize / 2
				checkv1x.Y -= tilesize / 2
				checkv2x := checkv1x
				checkv2x.X += tilesize
				checkv3x := checkv2x
				checkv3x.Y += tilesize
				checkv4x := checkv3x
				checkv4x.X -= tilesize

				canmove1, canmove2, canmove3, canmove4 := false, false, false, false

				checkobjs := false

				for a := 0; a < len(visroom); a++ {

					for b := 0; b < len(visroom[a].roomrec); b++ {

						if rl.CheckCollisionPointRec(checkv1x, visroom[a].roomrec[b].rec) {
							canmove1 = true
						}
						if rl.CheckCollisionPointRec(checkv2x, visroom[a].roomrec[b].rec) {
							canmove2 = true
						}
						if rl.CheckCollisionPointRec(checkv3x, visroom[a].roomrec[b].rec) {
							canmove3 = true
						}
						if rl.CheckCollisionPointRec(checkv4x, visroom[a].roomrec[b].rec) {
							canmove4 = true
						}

						if len(visroom[a].objs) > 0 {
							checkobjs = true
						}

					}
				}

				if checkobjs {
					for a := 0; a < len(visroom); a++ {
						if len(visroom[a].objs) > 0 {

							for b := 0; b < len(visroom[a].objs); b++ {

								if rl.CheckCollisionPointRec(checkv1x, visroom[a].objs[b].rec) && visroom[a].objs[b].solid {
									canmove1 = false
								}
								if rl.CheckCollisionPointRec(checkv2x, visroom[a].objs[b].rec) && visroom[a].objs[b].solid {
									canmove2 = false
								}
								if rl.CheckCollisionPointRec(checkv3x, visroom[a].objs[b].rec) && visroom[a].objs[b].solid {
									canmove3 = false
								}
								if rl.CheckCollisionPointRec(checkv4x, visroom[a].objs[b].rec) && visroom[a].objs[b].solid {
									canmove4 = false
								}

							}
						}

					}
				}

				if canmove1 && canmove2 && canmove3 && canmove4 {
					boss[a].cnt.X += boss[a].dirx
					boss[a].rec.X += boss[a].dirx
				}

				checkv1y := boss[a].cnt
				checkv1y.Y += boss[a].diry

				checkv1y.X -= tilesize / 2
				checkv1y.Y -= tilesize / 2
				checkv2y := checkv1y
				checkv2y.X += tilesize
				checkv3y := checkv2y
				checkv3y.Y += tilesize
				checkv4y := checkv3y
				checkv4y.X -= tilesize

				canmove1, canmove2, canmove3, canmove4 = false, false, false, false

				for a := 0; a < len(visroom); a++ {

					for b := 0; b < len(visroom[a].roomrec); b++ {

						if rl.CheckCollisionPointRec(checkv1y, visroom[a].roomrec[b].rec) {
							canmove1 = true
						}
						if rl.CheckCollisionPointRec(checkv2y, visroom[a].roomrec[b].rec) {
							canmove2 = true
						}
						if rl.CheckCollisionPointRec(checkv3y, visroom[a].roomrec[b].rec) {
							canmove3 = true
						}
						if rl.CheckCollisionPointRec(checkv4y, visroom[a].roomrec[b].rec) {
							canmove4 = true
						}

					}
				}

				if checkobjs {
					for a := 0; a < len(visroom); a++ {
						if len(visroom[a].objs) > 0 {

							for b := 0; b < len(visroom[a].objs); b++ {

								if rl.CheckCollisionPointRec(checkv1y, visroom[a].objs[b].rec) && visroom[a].objs[b].solid {
									canmove1 = false
								}
								if rl.CheckCollisionPointRec(checkv2y, visroom[a].objs[b].rec) && visroom[a].objs[b].solid {
									canmove2 = false
								}
								if rl.CheckCollisionPointRec(checkv3y, visroom[a].objs[b].rec) && visroom[a].objs[b].solid {
									canmove3 = false
								}
								if rl.CheckCollisionPointRec(checkv1y, visroom[a].objs[b].rec) && visroom[a].objs[b].solid {
									canmove4 = false
								}

							}
						}

					}
				}

				if canmove1 && canmove2 && canmove3 && canmove4 {
					boss[a].cnt.Y += boss[a].diry
					boss[a].rec.Y += boss[a].diry
					if boss[a].attacktype == 4 {
						count := 0
						bulletnum := 0
						if len(enemybullets) > 0 {
							for a := 0; a < len(enemybullets); a++ {
								if enemybullets[a].bossbullet {
									count++
									bulletnum = a
								}
							}

							if count == 1 {
								enemybullets[bulletnum].angle = angle2points(player.cnt, enemybullets[bulletnum].cnt)
							}

						}

					}
				}

				if rl.CheckCollisionPointRec(boss[a].cnt, player.boundrec) {
					boss[a].followtimer = fps
					boss[a].follow = false
				}

			}

			switch boss[a].attacktype {

			case 5:
				if frames%60 == 0 {
					if boss[a].onoff {
						boss[a].onoff = false
					} else {
						boss[a].onoff = true
					}

					if boss[a].onoff {
						num := rInt(10, 20)

						for {
							zobj := xobj{}
							zobj.noimg = true
							zobj.color = brightyellow()
							zobj.color2 = brightyellow()
							zobj.bulletsize = tilesize / 2
							zobj.rad = tilesize / 8
							zobj.cnt = boss[a].cnt
							zobj.dirx = rFloat32(-5, -5)
							zobj.dirx = rFloat32(-5, -5)
							zobj.bossbullet = true
							zobj.atk = boss[a].atk

							enemybullets = append(enemybullets, zobj)

							num--
							if num == 0 {
								break
							}
						}
						boss[a].onoff = false
					}

				}

			case 4: //orbit

				if len(enemybullets) < 50 {
					if frames%60 == 0 {

						zobj := xobj{}
						zobj.img = teddyimg
						zobj.bulletsize = tilesize
						zobj.monsternum = a
						zobj.atk = boss[a].atk

						zobj.v1 = boss[a].cnt
						zobj.cnt = boss[a].cnt
						zobj.cnt.Y -= tilesize * 4

						zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)

						zobj.color = randomyellow()
						zobj.orbits = true
						zobj.bossbullet = true
						zobj.angle = angle2points(boss[a].cnt, zobj.cnt)

						enemybullets = append(enemybullets, zobj)

						boss[a].onoff = true

					}
				}

			case 3: //laser
				if frames%60 == 0 {
					if boss[a].onoff {
						boss[a].onoff = false
					} else {
						if len(enemybullets) >= 1 {
							for b := 0; b < len(enemybullets); b++ {
								if enemybullets[b].bossbullet {
									enemybullets = remobj(enemybullets, b)
								}
							}
						}
						boss[a].onoff = true
					}

					if boss[a].onoff {

						zobj := xobj{}
						zobj.noimg = true
						zobj.line = true
						zobj.v1 = boss[a].cnt
						zobj.v2 = player.cnt
						zobj.color = randomgreen()
						zobj.atk = boss[a].atk
						zobj.monsternum = a
						zobj.bossbullet = true

						enemybullets = append(enemybullets, zobj)

						boss[a].onoff = false
					}
				}
			case 2: //bomb
				if frames%30 == 0 {
					zobj := xobj{}
					zobj.img = bombimg
					zobj.color = rl.White
					zobj.v1 = findrandpointinroom(player.cnt)
					dirx, diry := getpointdirs(boss[a].cnt, zobj.v1, 5)
					zobj.dirx = dirx
					zobj.diry = diry
					zobj.atk = boss[a].atk

					zobj.stops = true
					zobj.bossbullet = true

					zobj.cnt = boss[a].cnt
					zobj.bulletsize = tilesize

					zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)

					zobj.endrec = rl.NewRectangle(zobj.v1.X-tilesize/2, zobj.v1.Y-tilesize/2, tilesize, tilesize)

					enemybullets = append(enemybullets, zobj)

				}
			case 1: // missile
				if frames%boss[a].atkspeed == 0 {
					zobj := xobj{}
					zobj.img = missileimg
					zobj.color = brightred()
					dirx, diry := getpointdirs(boss[a].cnt, player.cnt, 4)
					zobj.ro = angle2points(boss[a].cnt, player.cnt) - 45
					zobj.dirx = dirx
					zobj.diry = diry
					zobj.atk = boss[a].atk
					zobj.bossbullet = true
					zobj.cnt = boss[a].cnt
					zobj.bulletsize = tilesize

					zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
					enemybullets = append(enemybullets, zobj)

				}

			case 0: // random direction bullets
				if frames%boss[a].atkspeed == 0 {
					zobj := xobj{}
					zobj.img = boss[a].bulletimg
					zobj.color = randomorange()
					v2 := boss[a].cnt
					v2.X += rFloat32(-10, 10)
					v2.Y += rFloat32(-10, 10)
					dirx, diry := getpointdirs(boss[a].cnt, v2, 4)
					zobj.dirx = dirx
					zobj.diry = diry
					zobj.atk = boss[a].atk
					zobj.rotates = true
					zobj.cnt = boss[a].cnt
					zobj.bulletsize = tilesize
					zobj.bossbullet = true
					zobj.grows = true
					zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
					enemybullets = append(enemybullets, zobj)

				}

			}
		}
	}

}
func upgemvalues() { //MARK: upgemvalues

	invenamount := 0

	beltonly := true

	for a := 0; a < len(inven); a++ {

		if inven[a].name == "gem" {
			invenamount += inven[a].amount
			beltonly = false
		}
	}

	for a := 0; a < len(beltinven); a++ {
		if beltinven[a].name == "gem" {
			invenamount += beltinven[a].amount
		}
	}

	if beltonly {
		var invenum []int
		if invenamount != gemstotal {
			for a := 0; a < len(beltinven); a++ {
				if beltinven[a].name == "gem" {
					invenum = append(invenum, a)
				}
			}

			if len(invenum) == 1 {
				beltinven[invenum[0]].amount = gemstotal
				beltinven[invenum[0]].name2 = "gem value is " + fmt.Sprint(beltinven[invenum[0]].amount)
				beltinven[invenum[0]].ident = true
			} else {

				for a := 1; a < len(invenum); a++ {
					beltinven[invenum[a]] = xobj{}
					findbeltinvennum()
				}
				beltinven[invenum[0]].amount = gemstotal
				beltinven[invenum[0]].name2 = "gem value is " + fmt.Sprint(beltinven[invenum[0]].amount)
				beltinven[invenum[0]].ident = true
			}

		}

	} else {

		var invenum []int

		if invenamount != gemstotal {
			for a := 0; a < len(inven); a++ {
				if inven[a].name == "gem" {
					invenum = append(invenum, a)
				}
			}

			if len(invenum) == 1 {
				inven[invenum[0]].amount = gemstotal
				inven[invenum[0]].name2 = "gem value is " + fmt.Sprint(inven[invenum[0]].amount)
				inven[invenum[0]].ident = true
			} else {

				for a := 1; a < len(invenum); a++ {
					inven[invenum[a]] = xobj{}
					findinvennum()
				}
				for a := 0; a < len(beltinven); a++ {

					if beltinven[a].kind == "gem" {
						beltinven[a] = xobj{}
						findbeltinvennum()
					}

				}

				inven[invenum[0]].amount = gemstotal
				inven[invenum[0]].name2 = "gem value is " + fmt.Sprint(inven[invenum[0]].amount)
				inven[invenum[0]].ident = true
			}
		}
	}

}
func uparmorset() { //MARK: uparmorset
	num := 0
	for a := 0; a < len(armorsetcount); a++ {

		if armorsetcount[a] > 3 {
			num = a
			break
		}
	}
	num++

	switch num {
	case 1:
		player.hpmax += 50
		if player.hpmax > 99 {
			player.hpmax = 99
		}
		newmsg("armor set of health equipped +50 MAX HP")
	case 2:
		player.str += 10
		if player.str > 99 {
			player.str = 99
		}
		player.intel += 10
		if player.intel > 99 {
			player.intel = 99
		}
		player.dex += 10
		if player.dex > 99 {
			player.dex = 99
		}
		player.luck += 10
		if player.luck > 99 {
			player.luck = 99
		}
		newmsg("armor set of boosting equipped +10 STR INT DEX LUK")

	case 3:
		player.luckorig = player.luck
		player.luck = 99
		newmsg("lucky luke armor set equipped 99 LUK")

	case 4:
		player.dexorig = player.dex
		player.dex = 99
		newmsg("robin hood armor set equipped 99 DEX")

	case 5:
		player.strorig = player.str
		player.str = 99
		newmsg("the hulk armor set equipped 99 STR")

	case 6:
		player.intelorig = player.intel
		player.intel = 99
		newmsg("merlin armor set equipped 99 INT")
	}
}
func uplegendary() { //MARK: uplegendary

	for a := 0; a < len(inven); a++ {

		if inven[a].legendary && inven[a].invenselect {

			switch inven[a].ability {
			case 8:
				for a := 0; a < len(inven); a++ {
					inven[a].ident = true
				}
				for a := 0; a < len(beltinven); a++ {
					beltinven[a].ident = true
				}
			case 6:
				if inven[a].timer > 0 {
					inven[a].timer--
				} else {
					activscroll.scrollnum = 1
					zmagic := xmagic{}
					magictimer = fps * 2
					zcircle := xcircle{}

					zcircle.rad = scrhf32 / 2
					zcircle.color = randomorange()
					zmagic.circles = append(zmagic.circles, zcircle)
					zcircle.rad -= 100
					zmagic.circles = append(zmagic.circles, zcircle)
					zcircle.rad -= 100
					zmagic.circles = append(zmagic.circles, zcircle)

					activmagic = append(activmagic, zmagic)
					magicon = true

					player.cnt, _ = findcntr()
					selpoint = player.cnt
					selpoint.X += tilesize / 4

					selrec = rl.NewRectangle(selpoint.X-tilesize/8, selpoint.Y-tilesize/8, tilesize/4, tilesize/4)
					player.dirx = 0
					player.diry = 0
					upplayer()

					inven[a].timer = rInt32(int(fps*10), int(fps*20))

				}
			case 1:
				if player.hp < player.hpmax {
					if frames%int(fps*4) == 0 {
						player.hp++
					}
				}

			}
		}

	}

}
func uppets() { //MARK: uppets

	for a := 0; a < len(pets); a++ {

		//movement
		if pets[a].timer > 0 {
			pets[a].timer--
		} else {
			if pets[a].idle {
				pets[a].idle = false
				pets[a].dirx = rFloat32(-pets[a].vel, pets[a].vel)
				pets[a].diry = rFloat32(-pets[a].vel, pets[a].vel)
				pets[a].timer = rInt(60, 120)
			} else {
				pets[a].idle = true
				pets[a].timer = rInt(60, 120)
			}
		}

		if !pets[a].idle {

			checkv2 := rl.NewVector2(pets[a].cnt.X+(pets[a].dirx*4), pets[a].cnt.Y+(pets[a].diry*4))

			if checkreccollisions(checkv2) {

				pets[a].cnt.X += pets[a].dirx
				pets[a].cnt.Y += pets[a].diry

				pets[a].rec = rl.NewRectangle(pets[a].cnt.X-tilesize/2, pets[a].cnt.Y-tilesize/2, tilesize, tilesize)
			} else {
				pets[a].idle = true
				pets[a].timer = rInt(60, 120)
			}

			if absdiff32(player.cnt.X, pets[a].cnt.X) > scrwf32/2 {

				if player.cnt.X > pets[a].cnt.X {
					pets[a].dirx = rFloat32(1, pets[a].vel)
				} else {
					pets[a].dirx = rFloat32(-pets[a].vel, -1)
				}
			}
			if absdiff32(player.cnt.Y, pets[a].cnt.Y) > scrhf32/2 {

				if player.cnt.Y > pets[a].cnt.Y {
					pets[a].diry = rFloat32(1, pets[a].vel)
				} else {
					pets[a].diry = rFloat32(-pets[a].vel, -1)
				}
			}
		}

		//offscreen
		if rl.CheckCollisionPointRec(pets[a].cnt, borderrec) {
			pets[a].offscreen = false
			pets[a].offscreenswitch = false

		} else {
			if !pets[a].offscreenswitch {
				pets[a].offscreen = true
				pets[a].offscreentimer = int(fps) * 5
				pets[a].offscreenswitch = true
			}
		}
		if pets[a].offscreentimer > 0 {
			pets[a].offscreentimer--
		}
		if pets[a].offscreentimer == 1 {

			pets[a].cnt = findrandpointinroom(player.cnt)
			pets[a].rec = rl.NewRectangle(pets[a].cnt.X-tilesize/2, pets[a].cnt.Y-tilesize/2, tilesize, tilesize)
			pets[a].idle = true
		}

	}

}
func upresistances() { //MARK: upresistances
	player.fireresistance = player.fireresistancejewel + player.fireresistancepotion
	player.poisonresistance = player.poisonresistancejewel + player.poisonresistancepotion
}
func uppotions() { //MARK: uppotions

	for a := 0; a < len(activpotions); a++ {

		if !activpotions[a].inactiv {
			if activpotions[a].timer > 0 {
				activpotions[a].timer--
			} else {
				switch activpotions[a].name2 {
				case "resist fire potion":
					player.fireresistance -= activpotions[a].amount
					activpotions[a].inactiv = true
				case "resist poison potion":
					player.poisonresistance -= activpotions[a].amount
					activpotions[a].inactiv = true
				}

			}

		}

	}

}

func upcams() { //MARK: upcams
	camera.Target = player.cnt
	camera.Offset.X = scrwf32 / 2
	camera.Offset.Y = scrhf32 / 2
}
func upplayer() { //MARK: upplayer

	if switchinvincible && !invinciblemode {
		invinciblemode = true
	} else if !switchinvincible && invinciblemode {
		invinciblemode = false
		player.hp = player.hpmax
	}

	if invinciblemode {
		player.hp = 99
	}

	if deathon {
		if player.hp <= 0 && !died {
			died = true
			pause = true
		}
	}

	if player.nomove {

		if camera.Zoom < 10 && !springswitch {
			camera.Zoom += 0.3
		} else {
			player.cnt = springv2

			selpoint = player.cnt
			selpoint.X += tilesize / 4

			selrec = rl.NewRectangle(selpoint.X-tilesize/8, selpoint.Y-tilesize/8, tilesize/4, tilesize/4)
			player.dirx = 0
			player.diry = 0

			player.rec = rl.NewRectangle(player.cnt.X-tilesize/2, player.cnt.Y-tilesize/2, tilesize, tilesize)
			player.boundrec.X = player.rec.X - tilesize*2
			player.boundrec.Y = player.rec.Y - tilesize*2
			borderrec = rl.NewRectangle(player.cnt.X-defw/2, player.cnt.Y-defh/2, defw, defh)
			visiblerec = borderrec
			visiblerec.X += invenrec.Width
			visiblerec.Y += msgrec.Height

			visiblerec.Width -= invenrec.Width
			visiblerec.Width -= statsrec.Width

			visiblerec.Height -= msgrec.Height
			visiblerec.Height -= footerrec.Height

			springswitch = true
		}

		if springswitch {
			if camera.Zoom > 1.5 {
				camera.Zoom -= 0.3

			} else {
				camera.Zoom = 1.5
				player.nomove = false
				springswitch = false
			}
		}

	} else {

		if selpoint != blankv2 && !rl.CheckCollisionPointRec(player.cnt, selrec) {

			animateplayer = true
			playeremoteon = false

			if selpoint.X > player.cnt.X {
				player.lr = true
			} else {
				player.lr = false
			}

			xdiff := absdiff32(player.cnt.X, selpoint.X)
			ydiff := absdiff32(player.cnt.Y, selpoint.Y)

			if xdiff > ydiff {
				player.dirx = player.vel
				if selpoint.X < player.cnt.X {
					player.dirx = -player.dirx
				}
				ychange := xdiff / player.vel
				player.diry = ydiff / ychange
				if selpoint.Y < player.cnt.Y {
					player.diry = -player.diry
				}
			} else {
				player.diry = player.vel
				if selpoint.Y < player.cnt.Y {
					player.diry = -player.diry
				}
				xchange := ydiff / player.vel
				player.dirx = xdiff / xchange
				if selpoint.X < player.cnt.X {
					player.dirx = -player.dirx
				}

			}

			checkv2x := player.cnt
			checkv2x.X += player.dirx
			player.v1 = checkv2x
			player.v1.X -= tilesize / 2
			player.v1.Y -= tilesize / 2
			player.v2 = player.v1
			player.v2.X += tilesize
			player.v3 = player.v2
			player.v3.Y += tilesize
			player.v4 = player.v3
			player.v4.X -= tilesize

			canmove1, canmove2, canmove3, canmove4 := false, false, false, false

			checkobjs := false

			for a := 0; a < len(visroom); a++ {

				for b := 0; b < len(visroom[a].roomrec); b++ {

					if rl.CheckCollisionPointRec(player.v1, visroom[a].roomrec[b].rec) {
						canmove1 = true
					}
					if rl.CheckCollisionPointRec(player.v2, visroom[a].roomrec[b].rec) {
						canmove2 = true
					}
					if rl.CheckCollisionPointRec(player.v3, visroom[a].roomrec[b].rec) {
						canmove3 = true
					}
					if rl.CheckCollisionPointRec(player.v4, visroom[a].roomrec[b].rec) {
						canmove4 = true
					}

					if len(visroom[a].objs) > 0 {
						checkobjs = true
					}

				}
			}

			if checkobjs {
				for a := 0; a < len(visroom); a++ {
					if len(visroom[a].objs) > 0 {

						for b := 0; b < len(visroom[a].objs); b++ {

							if rl.CheckCollisionPointRec(player.v1, visroom[a].objs[b].rec) && visroom[a].objs[b].solid {
								canmove1 = false
							}
							if rl.CheckCollisionPointRec(player.v2, visroom[a].objs[b].rec) && visroom[a].objs[b].solid {
								canmove2 = false
							}
							if rl.CheckCollisionPointRec(player.v3, visroom[a].objs[b].rec) && visroom[a].objs[b].solid {
								canmove3 = false
							}
							if rl.CheckCollisionPointRec(player.v4, visroom[a].objs[b].rec) && visroom[a].objs[b].solid {
								canmove4 = false
							}

						}
					}

				}
			}

			if canmove1 && canmove2 && canmove3 && canmove4 {
				player.cnt.X += player.dirx
			}

			checkv2y := player.cnt
			checkv2y.Y += player.diry
			player.v1 = checkv2y
			player.v1.X -= tilesize / 2
			player.v1.Y -= tilesize / 2
			player.v2 = player.v1
			player.v2.X += tilesize
			player.v3 = player.v2
			player.v3.Y += tilesize
			player.v4 = player.v3
			player.v4.X -= tilesize

			canmove1, canmove2, canmove3, canmove4 = false, false, false, false

			for a := 0; a < len(visroom); a++ {

				for b := 0; b < len(visroom[a].roomrec); b++ {

					if rl.CheckCollisionPointRec(player.v1, visroom[a].roomrec[b].rec) {
						canmove1 = true
					}
					if rl.CheckCollisionPointRec(player.v2, visroom[a].roomrec[b].rec) {
						canmove2 = true
					}
					if rl.CheckCollisionPointRec(player.v3, visroom[a].roomrec[b].rec) {
						canmove3 = true
					}
					if rl.CheckCollisionPointRec(player.v4, visroom[a].roomrec[b].rec) {
						canmove4 = true
					}

				}
			}

			if checkobjs {
				for a := 0; a < len(visroom); a++ {
					if len(visroom[a].objs) > 0 {

						for b := 0; b < len(visroom[a].objs); b++ {

							if rl.CheckCollisionPointRec(player.v1, visroom[a].objs[b].rec) && visroom[a].objs[b].solid {
								canmove1 = false
							}
							if rl.CheckCollisionPointRec(player.v2, visroom[a].objs[b].rec) && visroom[a].objs[b].solid {
								canmove2 = false
							}
							if rl.CheckCollisionPointRec(player.v3, visroom[a].objs[b].rec) && visroom[a].objs[b].solid {
								canmove3 = false
							}
							if rl.CheckCollisionPointRec(player.v4, visroom[a].objs[b].rec) && visroom[a].objs[b].solid {
								canmove4 = false
							}

						}
					}

				}
			}

			if canmove1 && canmove2 && canmove3 && canmove4 {
				player.cnt.Y += player.diry
			}

			player.v1 = player.cnt
			player.v1.X -= tilesize / 2
			player.v1.Y -= tilesize / 2
			player.v2 = player.v1
			player.v2.X += tilesize
			player.v3 = player.v2
			player.v3.Y += tilesize
			player.v4 = player.v3
			player.v4.X -= tilesize

			if player.flametrail {

				// add flames
				if rolldice() > 4 {
					zfx := xfx{}
					v2 := (player.cnt)
					v2.Y -= tilesize / 2
					zfx.name = "flame"
					zfx.img = flameimg
					zfx.rec = rl.NewRectangle(v2.X, v2.Y, tilesize, tilesize)
					zfx.timer = rInt32(fpsint, fpsint*3)
					fx = append(fx, zfx)
				}

			}

		} else {
			animateplayer = false
			playerimg.X = 1
			playerlimg.X = 66
		}

		player.rec = rl.NewRectangle(player.cnt.X-tilesize/2, player.cnt.Y-tilesize/2, tilesize, tilesize)
		player.boundrec.X = player.rec.X - tilesize*2
		player.boundrec.Y = player.rec.Y - tilesize*2

		if player.weapon.name != "" {
			player.weapon.meleerangerec.X = player.cnt.X - player.weapon.meleerangerec.Width/2
			player.weapon.meleerangerec.Y = player.cnt.Y - player.weapon.meleerangerec.Height/2
		}

		borderrec = rl.NewRectangle(player.cnt.X-defw/2, player.cnt.Y-defh/2, defw, defh)
		visiblerec = borderrec
		visiblerec.X += invenrec.Width
		visiblerec.Y += msgrec.Height

		visiblerec.Width -= invenrec.Width
		visiblerec.Width -= statsrec.Width

		visiblerec.Height -= msgrec.Height
		visiblerec.Height -= footerrec.Height
	}
}
func upjewel(onoff bool, jewel xobj) { //MARK: upjewel

	if onoff {
		switch jewel.usetype {
		case 1:
			player.dex += jewel.amount
			newmsg("dexterity increased by " + fmt.Sprint(jewel.amount))
		case 2:
			player.str += jewel.amount
			newmsg("strength increased by " + fmt.Sprint(jewel.amount))
		case 3:
			player.luck += jewel.amount
			newmsg("luck increased by " + fmt.Sprint(jewel.amount))
		case 4:
			player.intel += jewel.amount
			newmsg("intelligence increased by " + fmt.Sprint(jewel.amount))
		case 5:
			player.fireresistancejewel += jewel.amount * 10
			upresistances()
			txt := fmt.Sprint(player.fireresistance)
			if player.fireresistance > 100 {
				txt = "100"
			}
			newmsg("fire resistance increased - " + txt + "%" + " chance of not catching fire")
		case 6:
			player.poisonresistancejewel += jewel.amount * 10
			upresistances()
			txt := fmt.Sprint(player.poisonresistance)
			if player.poisonresistance > 100 {
				txt = "100"
			}
			newmsg("poison resistance increased - " + txt + "%" + " chance of not getting poisoned")
		case 7:
			player.immune = true
			player.sick = false
			player.sicktimer = 0
			newmsg("you are now immune to disease")
		case 8:
			player.hpmax += jewel.amount
			newmsg("max hp increased by " + fmt.Sprint(jewel.amount))
		}
	} else {

		switch jewel.usetype {
		case 1:
			player.dex -= jewel.amount
		case 2:
			player.str -= jewel.amount
		case 3:
			player.luck -= jewel.amount
		case 4:
			player.intel -= jewel.amount
		case 5:
			player.fireresistancejewel -= jewel.amount * 10
			upresistances()
			txt := fmt.Sprint(player.fireresistance)
			if player.fireresistance > 100 {
				txt = "100"
			} else if player.fireresistance < 100 {
				newmsg("fire resistance decreased - " + txt + "%" + " chance of not catching fire")
			}
		case 6:
			player.poisonresistancejewel -= jewel.amount * 10
			upresistances()
			txt := fmt.Sprint(player.poisonresistance)
			if player.poisonresistance > 100 {
				txt = "100"
			} else if player.poisonresistance < 100 {
				newmsg("poison resistance decreased - " + txt + "%" + " chance of not getting poisoned")
			}
		case 7:
			player.immune = false
			newmsg("you are no longer immune to disease")
		case 8:
			player.hpmax -= jewel.amount
			newmsg("max hp decreased by " + fmt.Sprint(jewel.amount))
		}

	}

}
func upjewelteleport() { //MARK: upjewelteleport

	for a := 0; a < len(inven); a++ {
		if inven[a].kind == "jewel" && inven[a].usetype == 9 && inven[a].invenselect {
			inven[a].timer--
			if inven[a].timer <= 0 {
				activscroll.scrollnum = 1
				zmagic := xmagic{}
				magictimer = fps * 2
				zcircle := xcircle{}

				zcircle.rad = scrhf32 / 2
				zcircle.color = randomorange()
				zmagic.circles = append(zmagic.circles, zcircle)
				zcircle.rad -= 100
				zmagic.circles = append(zmagic.circles, zcircle)
				zcircle.rad -= 100
				zmagic.circles = append(zmagic.circles, zcircle)

				activmagic = append(activmagic, zmagic)
				magicon = true

				player.cnt, _ = findcntr()
				selpoint = player.cnt
				selpoint.X += tilesize / 4

				selrec = rl.NewRectangle(selpoint.X-tilesize/8, selpoint.Y-tilesize/8, tilesize/4, tilesize/4)
				player.dirx = 0
				player.diry = 0
				upplayer()
				inven[a].timer = rInt32(fpsint*5, fpsint*10)
			}

		}
	}

}
func uprooms() { //MARK: uprooms

	visroom = nil

	for a := 0; a < len(level); a++ {
		if rl.CheckCollisionRecs(level[a].boundrec, borderrec) {
			level[a].vis = true

			visroom = append(visroom, level[a])
		} else {
			level[a].vis = false
		}
	}

}

func upmonsters() { //MARK: upmonsters

	vismonsters = nil

	for a := 0; a < len(monsters); a++ {
		if !monsters[a].inactiv {
			//move
			switch monsters[a].move {
			case 4: //random xy
				switch monsters[a].movenum {
				case 0:
					checkv2 := monsters[a].cnt
					checkv2.Y += tilesize
					checkv2.X += tilesize

					monstervel := rFloat32(0.5, 4)
					newv2 := monsters[a].cnt
					newv2.Y += monstervel
					monstervel = rFloat32(0.5, 4)
					newv2.X += monstervel

					if checkreccollisions(checkv2) {
						monsters[a].cnt = newv2
					} else {
						monsters[a].movenum = rInt(0, 8)
					}

					if rolldice()+rolldice() == 12 {
						monsters[a].movenum = rInt(0, 8)
					}
				case 1:
					checkv2 := monsters[a].cnt
					checkv2.Y -= tilesize
					checkv2.X += tilesize

					monstervel := rFloat32(0.5, 4)
					newv2 := monsters[a].cnt
					newv2.Y -= monstervel
					monstervel = rFloat32(0.5, 4)
					newv2.X += monstervel

					if checkreccollisions(checkv2) {
						monsters[a].cnt = newv2
					} else {
						monsters[a].movenum = rInt(0, 8)
					}

					if rolldice()+rolldice() == 12 {
						monsters[a].movenum = rInt(0, 8)
					}

				case 2:
					checkv2 := monsters[a].cnt
					checkv2.Y -= tilesize
					checkv2.X -= tilesize

					monstervel := rFloat32(0.5, 4)
					newv2 := monsters[a].cnt
					newv2.Y -= monstervel
					monstervel = rFloat32(0.5, 4)
					newv2.X -= monstervel

					if checkreccollisions(checkv2) {
						monsters[a].cnt = newv2
					} else {
						monsters[a].movenum = rInt(0, 8)
					}

					if rolldice()+rolldice() == 12 {
						monsters[a].movenum = rInt(0, 8)
					}

				case 3:
					checkv2 := monsters[a].cnt
					checkv2.Y += tilesize
					checkv2.X -= tilesize

					monstervel := rFloat32(0.5, 4)
					newv2 := monsters[a].cnt
					newv2.Y += monstervel
					monstervel = rFloat32(0.5, 4)
					newv2.X -= monstervel

					if checkreccollisions(checkv2) {
						monsters[a].cnt = newv2
					} else {
						monsters[a].movenum = rInt(0, 8)
					}

					if rolldice()+rolldice() == 12 {
						monsters[a].movenum = rInt(0, 8)
					}

				case 4:
					checkv2 := monsters[a].cnt
					checkv2.Y += tilesize

					monstervel := rFloat32(0.5, 4)
					newv2 := monsters[a].cnt
					newv2.Y += monstervel

					if checkreccollisions(checkv2) {
						monsters[a].cnt = newv2
					} else {
						monsters[a].movenum = rInt(0, 8)
					}

					if rolldice()+rolldice() == 12 {
						monsters[a].movenum = rInt(0, 8)
					}

				case 5:
					checkv2 := monsters[a].cnt
					checkv2.Y -= tilesize

					monstervel := rFloat32(0.5, 4)
					newv2 := monsters[a].cnt
					newv2.Y -= monstervel

					if checkreccollisions(checkv2) {
						monsters[a].cnt = newv2
					} else {
						monsters[a].movenum = rInt(0, 8)
					}

					if rolldice()+rolldice() == 12 {
						monsters[a].movenum = rInt(0, 8)
					}

				case 6:
					checkv2 := monsters[a].cnt
					checkv2.X -= tilesize

					monstervel := rFloat32(0.5, 4)
					newv2 := monsters[a].cnt
					newv2.X -= monstervel

					if checkreccollisions(checkv2) {
						monsters[a].cnt = newv2
					} else {
						monsters[a].movenum = rInt(0, 8)
					}

					if rolldice()+rolldice() == 12 {
						monsters[a].movenum = rInt(0, 8)
					}

				case 7:
					checkv2 := monsters[a].cnt
					checkv2.X += tilesize

					monstervel := rFloat32(0.5, 4)
					newv2 := monsters[a].cnt
					newv2.X += monstervel

					if checkreccollisions(checkv2) {
						monsters[a].cnt = newv2
					} else {
						monsters[a].movenum = rInt(0, 8)
					}

					if rolldice()+rolldice() == 12 {
						monsters[a].movenum = rInt(0, 8)
					}

				}

			case 3: // random

				switch monsters[a].movenum {
				case 0:
					checkv2 := monsters[a].cnt
					checkv2.Y += tilesize
					monstervel := rFloat32(0.5, 4)
					newv2 := monsters[a].cnt
					newv2.Y += monstervel
					if checkreccollisions(checkv2) {
						monsters[a].cnt = newv2
					} else {
						monsters[a].movenum = rInt(0, 8)
					}
				case 1:
					checkv2 := monsters[a].cnt
					checkv2.X += tilesize
					monstervel := rFloat32(0.5, 4)
					newv2 := monsters[a].cnt
					newv2.X += monstervel
					if checkreccollisions(checkv2) {
						monsters[a].cnt = newv2

					} else {
						monsters[a].movenum = rInt(0, 8)
					}
				case 2:
					checkv2 := monsters[a].cnt
					checkv2.Y -= tilesize
					monstervel := rFloat32(0.5, 4)
					newv2 := monsters[a].cnt
					newv2.Y -= monstervel
					if checkreccollisions(checkv2) {
						monsters[a].cnt = newv2

					} else {
						monsters[a].movenum = rInt(0, 8)
					}

				case 3:
					checkv2 := monsters[a].cnt
					checkv2.X -= tilesize
					monstervel := rFloat32(0.5, 4)
					newv2 := monsters[a].cnt
					newv2.X -= monstervel
					if checkreccollisions(checkv2) {
						monsters[a].cnt = newv2

					} else {
						monsters[a].movenum = rInt(0, 8)
					}

				case 4:
					checkv2 := monsters[a].cnt
					checkv2.X += tilesize
					checkv2.Y -= tilesize
					monstervel := rFloat32(0.5, 4)
					newv2 := monsters[a].cnt
					newv2.X += monstervel
					newv2.Y -= monstervel

					if checkreccollisions(checkv2) {
						monsters[a].cnt = newv2

					} else {
						monsters[a].movenum = rInt(0, 8)
					}

				case 5:
					checkv2 := monsters[a].cnt
					checkv2.X += tilesize
					checkv2.Y += tilesize
					monstervel := rFloat32(0.5, 4)
					newv2 := monsters[a].cnt
					newv2.X += monstervel
					newv2.Y += monstervel

					if checkreccollisions(checkv2) {
						monsters[a].cnt = newv2

					} else {
						monsters[a].movenum = rInt(0, 8)
					}

				case 6:
					checkv2 := monsters[a].cnt
					monstervel := rFloat32(0.5, 4)
					checkv2.X -= monstervel * 4
					checkv2.Y += monstervel * 4

					newv2 := monsters[a].cnt
					newv2.X -= monstervel
					newv2.Y += monstervel

					if checkreccollisions(checkv2) {
						monsters[a].cnt = newv2

					} else {
						monsters[a].movenum = rInt(0, 8)
					}
				case 7:
					checkv2 := monsters[a].cnt
					monstervel := rFloat32(0.5, 4)
					checkv2.X -= monstervel * 4
					checkv2.Y -= monstervel * 4

					newv2 := monsters[a].cnt
					newv2.X -= monstervel
					newv2.Y -= monstervel

					if checkreccollisions(checkv2) {
						monsters[a].cnt = newv2

					} else {
						monsters[a].movenum = rInt(0, 8)
					}
				}
				if rolldice()+rolldice() == 12 {
					monsters[a].movenum = rInt(0, 8)
				}
			case 2: // up down left right
				if monsters[a].moveswitch {
					if monsters[a].moveswitch2 {
						checkv2 := monsters[a].cnt
						checkv2.Y += tilesize
						monstervel := rFloat32(0.5, 4)
						newv2 := monsters[a].cnt
						newv2.Y += monstervel
						if checkreccollisions(checkv2) {
							monsters[a].cnt = newv2

						} else {
							monsters[a].moveswitch = false
						}

						if rolldice()+rolldice() == 12 {
							monsters[a].moveswitch2 = false
						}
					} else {
						checkv2 := monsters[a].cnt
						checkv2.X += tilesize
						monstervel := rFloat32(0.5, 4)
						newv2 := monsters[a].cnt
						newv2.X += monstervel
						if checkreccollisions(checkv2) {
							monsters[a].cnt = newv2

						} else {
							monsters[a].moveswitch = false
						}

						if rolldice()+rolldice() == 12 {
							monsters[a].moveswitch2 = true
						}
					}
				} else {

					if monsters[a].moveswitch2 {
						checkv2 := monsters[a].cnt
						checkv2.Y -= tilesize
						monstervel := rFloat32(0.5, 4)
						newv2 := monsters[a].cnt
						newv2.Y -= monstervel
						if checkreccollisions(checkv2) {
							monsters[a].cnt = newv2

						} else {
							monsters[a].moveswitch = false
						}
						if rolldice()+rolldice() == 12 {
							monsters[a].moveswitch2 = false
						}
					} else {
						checkv2 := monsters[a].cnt
						checkv2.X -= tilesize
						monstervel := rFloat32(0.5, 4)
						newv2 := monsters[a].cnt
						newv2.X -= monstervel
						if checkreccollisions(checkv2) {
							monsters[a].cnt = newv2

						} else {
							monsters[a].moveswitch = true
						}

						if rolldice()+rolldice() == 12 {
							monsters[a].moveswitch2 = true
						}
					}
				}

			case 1: //zigzag

				if monsters[a].moveswitch {

					if monsters[a].moveswitch2 {
						checkv2 := monsters[a].cnt
						checkv2.X += tilesize
						checkv2.Y -= tilesize
						monstervel := rFloat32(0.5, 4)
						newv2 := monsters[a].cnt
						newv2.X += monstervel
						newv2.Y -= monstervel

						if checkreccollisions(checkv2) {
							monsters[a].cnt = newv2

						} else {
							monsters[a].moveswitch = false
						}
						if rolldice()+rolldice() == 12 {
							monsters[a].moveswitch2 = false
						}
					} else {
						checkv2 := monsters[a].cnt
						checkv2.X += tilesize
						checkv2.Y += tilesize
						monstervel := rFloat32(0.5, 4)
						newv2 := monsters[a].cnt
						newv2.X += monstervel
						newv2.Y += monstervel

						if checkreccollisions(checkv2) {
							monsters[a].cnt = newv2

						} else {
							monsters[a].moveswitch = false
						}

						if rolldice()+rolldice() == 12 {
							monsters[a].moveswitch2 = true
						}
					}
				} else {

					if monsters[a].moveswitch2 {
						checkv2 := monsters[a].cnt
						monstervel := rFloat32(0.5, 4)
						checkv2.X -= monstervel * 4
						checkv2.Y += monstervel * 4

						newv2 := monsters[a].cnt
						newv2.X -= monstervel
						newv2.Y += monstervel

						if checkreccollisions(checkv2) {
							monsters[a].cnt = newv2

						} else {
							monsters[a].moveswitch = true
						}
						if rolldice()+rolldice() == 12 {
							monsters[a].moveswitch2 = false
						}
					} else {
						checkv2 := monsters[a].cnt
						monstervel := rFloat32(0.5, 4)
						checkv2.X -= monstervel * 4
						checkv2.Y -= monstervel * 4

						newv2 := monsters[a].cnt
						newv2.X -= monstervel
						newv2.Y -= monstervel

						if checkreccollisions(checkv2) {
							monsters[a].cnt = newv2

						} else {
							monsters[a].moveswitch = true
						}
						if rolldice()+rolldice() == 12 {
							monsters[a].moveswitch2 = true
						}
					}
				}
			}

			monsters[a].rec = rl.NewRectangle(monsters[a].cnt.X-tilesize/2, monsters[a].cnt.Y-tilesize/2, tilesize, tilesize)

			if len(fx) > 0 {

				for b := 0; b < len(fx); b++ {

					if rl.CheckCollisionRecs(monsters[a].rec, fx[b].rec) {

						if !monsters[a].burning {
							monsters[a].burning = true
							monsters[a].timer = 3 * fps
						}

					}

				}
			}

			switch monsters[a].atktype {

			case 4: //poison
				if rolldice()+rolldice() == 12 {
					zobj := xobj{}
					zobj.monsternum = a
					zobj.cnt = monsters[a].cnt
					zobj.dirx = rFloat32(-tilesize/5, tilesize/5)
					zobj.diry = rFloat32(-tilesize/5, tilesize/5)
					zobj.color = randomgreen()
					zobj.bulletsize = tilesize / 2
					zobj.atk = 1 + currentlevelnum
					zobj.img = fireballimg
					zobj.poison = true
					v2 := monsters[a].cnt
					v2.X += zobj.dirx * 4
					v2.Y += zobj.diry * 4
					zobj.ro = angle2points(monsters[a].cnt, v2) + 180
					zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/4, zobj.cnt.Y-tilesize/4, tilesize/2, tilesize/2)
					enemybullets = append(enemybullets, zobj)
				}
			case 3: //multi direction
				if rolldice()+rolldice()+rolldice() == 18 {

					dir := rFloat32(tilesize/8, tilesize/4)

					zobj := xobj{}
					zobj.monsternum = a
					zobj.noimg = true
					zobj.cnt = monsters[a].cnt

					zobj.color = randomorange()
					zobj.color2 = randomcolor()
					zobj.bulletsize = tilesize / 2
					zobj.atk = 1 + currentlevelnum
					zobj.rad = tilesize / 4
					zobj.color = randomcolor()
					zobj.color2 = randomcolor()
					zobj.bulletsize = tilesize / 2
					zobj.atk = 1 + currentlevelnum

					zobj.dirx = dir

					enemybullets = append(enemybullets, zobj)

					zobj.dirx = -dir
					enemybullets = append(enemybullets, zobj)
					zobj.diry = -dir
					zobj.dirx = 0
					enemybullets = append(enemybullets, zobj)
					zobj.diry = dir
					enemybullets = append(enemybullets, zobj)

				}

			case 2: //fireball
				if rolldice()+rolldice() == 12 {
					zobj := xobj{}
					zobj.monsternum = a
					zobj.cnt = monsters[a].cnt
					zobj.dirx = rFloat32(-tilesize/5, tilesize/5)
					zobj.diry = rFloat32(-tilesize/5, tilesize/5)
					zobj.color = randomorange()
					zobj.bulletsize = tilesize / 2
					zobj.atk = 1 + currentlevelnum
					zobj.img = fireballimg
					zobj.burns = true
					v2 := monsters[a].cnt
					v2.X += zobj.dirx * 4
					v2.Y += zobj.diry * 4
					zobj.ro = angle2points(monsters[a].cnt, v2) + 180
					zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/4, zobj.cnt.Y-tilesize/4, tilesize/2, tilesize/2)
					enemybullets = append(enemybullets, zobj)
				}

			case 1: //random direction
				if rolldice()+rolldice() == 12 {
					zobj := xobj{}
					zobj.monsternum = a
					zobj.noimg = true
					zobj.cnt = monsters[a].cnt
					zobj.dirx = rFloat32(-tilesize/5, tilesize/5)
					zobj.diry = rFloat32(-tilesize/5, tilesize/5)
					zobj.rad = tilesize / 4
					zobj.color = randomcolor()
					zobj.color2 = randomcolor()
					zobj.bulletsize = tilesize / 2
					zobj.atk = 1 + currentlevelnum

					zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/4, zobj.cnt.Y-tilesize/4, tilesize/2, tilesize/2)
					enemybullets = append(enemybullets, zobj)
				}

			}

			//animation
			if frames%monsteranimatetimer == 0 {
				monsters[a].img.X += monsters[a].rec.Width
				if monsters[a].img.X >= monsters[a].endx {
					monsters[a].img.X = monsters[a].startx
				}
			}

			//add to vismonsters
			if rl.CheckCollisionRecs(monsters[a].rec, borderrec) {
				vismonsters = append(vismonsters, monsters[a])
			}

			if monsters[a].hppause > 0 {
				monsters[a].hppause--
			}

		}

	}

}
func upweapons() { //MARK: upweapons

	clear := false

	for a := 0; a < len(activweapons); a++ {
		if activweapons[a].noimg {
			switch activweapons[a].usetype {

			case 3:
				if activweapons[a].timer > 0 {
					activweapons[a].timer--
				} else {
					activweapons[a].inactiv = true
					clear = true
				}

			case 2:

				activweapons[a].angle = activweapons[a].angle * (math.Pi / 180)

				newx := float32(math.Cos(float64(activweapons[a].angle)))*(activweapons[a].v1.X-activweapons[a].v2.X) - float32(math.Sin(float64(activweapons[a].angle)))*(activweapons[a].v1.Y-activweapons[a].v2.Y) + activweapons[a].v2.X

				newy := float32(math.Sin(float64(activweapons[a].angle)))*(activweapons[a].v1.X-activweapons[a].v2.X) + float32(math.Cos(float64(activweapons[a].angle)))*(activweapons[a].v1.Y-activweapons[a].v2.Y) + activweapons[a].v2.Y

				activweapons[a].v1 = rl.NewVector2(newx, newy)

				activweapons[a].angle += 2

				if activweapons[a].timer > 0 {
					activweapons[a].timer--
				} else {
					activweapons[a].inactiv = true
					clear = true
				}

			case 1:
				if activweapons[a].timer > 0 {
					activweapons[a].timer--
				} else {
					activweapons[a].inactiv = true
					clear = true
				}
			}

			checkweaponcollisions()
		} else {
			if !activweapons[a].inactiv {
				activweapons[a].rec.X += activweapons[a].dirx
				activweapons[a].rec.Y += activweapons[a].diry

				if activweapons[a].rec.X > borderrec.X+borderrec.Width {
					activweapons[a].inactiv = true
					clear = true
				}
				if activweapons[a].rec.X < borderrec.X {
					activweapons[a].inactiv = true
					clear = true
				}

				if activweapons[a].rec.Y > borderrec.Y+borderrec.Height {
					activweapons[a].inactiv = true
					clear = true
				}
				if activweapons[a].rec.Y < borderrec.Y {
					activweapons[a].inactiv = true
					clear = true
				}

				checkweaponcollisions()
			}
		}
	}

	if clear {
		for a := 0; a < len(activweapons); a++ {
			if activweapons[a].inactiv {
				activweapons = remobj(activweapons, a)
			}
		}
	}

}
func updedmonsters() { //MARK: updedmonsters

	clear := false
	for a := 0; a < len(dedmonsters); a++ {
		if !dedmonsters[a].inactiv {
			for b := 0; b < len(dedmonsters[a].circles); b++ {
				dedmonsters[a].circles[b].rad -= 2
				if dedmonsters[a].circles[b].rad <= 0 {
					dedmonsters[a].inactiv = true
					clear = true
				}
			}
		}
	}

	if clear {
		for a := 0; a < len(dedmonsters); a++ {
			if dedmonsters[a].inactiv {

				dedmonsters = remdedmons(dedmonsters, a)
			}
		}
	}

}

func upexplosions() { //MARK: upexplosions

	clear := false
	for a := 0; a < len(xplodecircs); a++ {
		if !xplodecircs[a].inactiv {
			if player.hpppause == 0 {
				if rl.CheckCollisionCircleRec(xplodecircs[a].v2, xplodecircs[a].rad, player.rec) {

					player.hp -= xplodecircs[a].atk
					if soundfxon && !mute {
						rl.PlaySoundMulti(playerdamageaud)
					}
					if player.hp < 0 {
						player.hp = 0
					}
					player.hpppause = fps

				}
			}

			xplodecircs[a].rad -= 2
			if xplodecircs[a].rad <= 0 {
				xplodecircs[a].inactiv = true
				clear = true
			}

		}
	}

	if clear {
		for a := 0; a < len(xplodecircs); a++ {
			if xplodecircs[a].inactiv {

				xplodecircs = remcirc(xplodecircs, a)
			}
		}
	}

}
func upruntime() { //MARK: upruntime

	secstotal := runtimer / 60

	mins, secs := secstotal/60, secstotal%60

	if mins > 60 {
		mins = 0
	}

	runmin = mins
	runsecs = secs

}

// MARK: ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC
func restart() { //MARK: restart

	pause = true
	currentlevelnum = 0
	invencurrentnum = 0
	beltinvencurrentnum = 0
	killcount = 0
	bosskills = 0
	monsterkills = 0

	player.coins = 0
	player.teleports = 0

	boss = nil
	pets = nil
	enemybullets = nil
	fx = nil

	for a := 0; a < len(inven); a++ {
		inven[a] = xobj{}
	}

	for a := 0; a < len(beltinven); a++ {
		beltinven[a] = xobj{}
	}

	runtimer = 0
	diedy = -scrhf32
	deathon = true

	introtxtx, introtxtx2 = -200, -700
	introtxty = -scrh / 3

	pigeononoff = true
	pigeontimer = fps * rInt32(20, 100)
	activammonum = blankint
	activweaponnum = blankint
	player.weapon = xobj{}
	player.ammo = xobj{}
	player.object = xobj{}
	player.poisoned = false
	player.poisonresistance = 0
	player.burning = false
	player.burntimer = 0
	player.sick = false
	player.sicktimer = 0

	makelevel()
	makeplayer()
	makepet()

	newmsg("so you have returned to try again, better luck this time")

	died = false
	scoreon = false

	pause = false

}

func explode(bulletnum int) { //MARK: explode

	num := rInt(12, 18)

	for {
		zcirc := xcircle{}
		zcirc.color = randomorange()
		zcirc.rad = rFloat32(60, 120)
		zcirc.v2 = enemybullets[bulletnum].cnt
		zcirc.v2.X += rFloat32(-tilesize*2, tilesize*2)
		zcirc.v2.Y += rFloat32(-tilesize*2, tilesize*2)
		zcirc.atk = enemybullets[bulletnum].atk

		xplodecircs = append(xplodecircs, zcirc)

		num--
		if num == 0 {
			break
		}

	}

	enemybullets[bulletnum].nodraw = true
	enemybullets[bulletnum].inactiv = true

}

func checkgemskeys() { //MARK: checkgemskeys

	gemstotal = 0
	keystotal = 0

	for a := 0; a < len(inven); a++ {

		if inven[a].name == "key" {
			keystotal++
		}
		if inven[a].name == "gem" {
			gemstotal += inven[a].amount
		}
	}

	for a := 0; a < len(beltinven); a++ {

		if beltinven[a].name == "key" {
			keystotal++
		}
		if beltinven[a].name == "gem" {
			gemstotal += inven[a].amount
		}
	}

}

func applyshopitem(itemnum int) { //MARK: applyshopitem

	switch shopitems[itemnum].name2 {

	case "armor":
		randomarmoron = true
	case "cactus":
		cactusimmunity = true
	case "coins":
		player.coins += shopitems[itemnum].amount
	case "dex":
		player.dex += shopitems[itemnum].amount
	case "disease":
	case "fireresist":
	case "firetrail":
	case "hp":
		player.hp = player.hpmax
	case "hpmax":
		player.hpmax += shopitems[itemnum].amount
		if player.hpmax > 99 {
			player.hpmax = 99
		}
	case "int":
		player.intel += shopitems[itemnum].amount
	case "luk":
		player.luck += shopitems[itemnum].amount
	case "poison":
	case "spike":
		spikeimmunity = true
	case "str":
		player.str += shopitems[itemnum].amount
	case "tel":
		player.teleports += shopitems[itemnum].amount
		if player.teleports > 99 {
			player.teleports = 99
		}
	case "tel99":
		player.teleports = 99

	}

}

func switchinven(objtype string, invennum int) { //MARK: switchinven

	for a := 0; a < len(inven); a++ {

		if a != invennum {
			if inven[a].legendary {
				if inven[a].name3 == objtype && inven[a].invenselect {
					inven[a].invenselect = false
					switch inven[invennum].ability {
					case 2:
						player.vampirelev--
					case 3:
						player.thornslev--
					}
				}

			} else {
				if inven[a].kind == objtype && inven[a].invenselect {
					inven[a].invenselect = false
				}
			}
		}

	}

}

func newmsg(txt string) { //MARK: newmsg

	msgs = append(msgs, txt)
	newmsgtimer = fps * 6
}
func choosesignmsg() string {

	signnames := []string{"My wife told me to stop impersonating a flamingo. I had to put my foot down.", "I went to buy some camo pants but couldn't find any.", "I failed math so many times at school, I can't even count.", "I used to have a handle on life, but then it broke.", "I was wondering why the frisbee kept getting bigger and bigger, but then it hit me.", "I heard there were a bunch of break-ins over at the car park. That is wrong on so many levels.", "When life gives you melons, you might be dyslexic.", "Don't you hate it when someone answers their own questions? I do.", "It takes a lot of balls to golf the way I do.", " I told him to be himself; that was pretty mean, I guess.", "I know they say that money talks, but all mine says is 'Goodbye.'", "My father has schizophrenia, but he's good people.", "The problem with kleptomaniacs is that they always take things literally.", "I can't believe I got fired from the calendar factory. All I did was take a day off.", "Most people are shocked when they find out how bad I am as an electrician.", "Never trust atoms; they make up everything.", "My wife just found out I replaced our bed with a trampoline. She hit the ceiling.", "I was addicted to the hokey pokey, but then I turned myself around.", "I used to think I was indecisive. But now I'm not so sure.", "Russian dolls are so full of themselves.", "The easiest time to add insult to injury is when you're signing someone's cast.", "Light travels faster than sound, which is the reason that some people appear bright before you hear them speak.", "My therapist says I have a preoccupation for revenge. We'll see about that.", "A termite walks into the bar and asks, 'Is the bar tender here?'", "A told my girlfriend she drew her eyebrows too high. She seemed surprised.", "People who use selfie sticks really need to have a good, long look at themselves.", "Two fish are in a tank. One says, 'How do you drive this thing?'", "I always take life with a grain of salt. And a slice of lemon. And a shot of tequila.", "Just burned 2,000 calories. That's the last time I leave brownies in the oven while I nap.", "Always borrow money from a pessimist. They'll never expect it back.", "Build a man a fire and he'll be warm for a day. Set a man on fire and he'll be warm for the rest of his life.", "I don't suffer from insanity—I enjoy every minute of it.", "The last thing I want to do is hurt you; but it's still on the list.", "The problem isn't that obesity runs in your family. It's that no one runs in your family.", "I'm reading a book about anti-gravity. It's impossible to put down.", "Atheism is a non-prophet organization.", "A recent study has found that women who carry a little extra weight live longer than the men who mention it.", "The future, the present, and the past walk into a bar. Things got a little tense.", "Last night my girlfriend was complaining that I never listen to her… or something like that.", "Maybe if we start telling people their brain is an app, they'll want to use it.", "If a parsley farmer gets sued, can they garnish his wages?", "I didn't think orthopedic shoes would help, but I stand corrected.", "People who take care of chickens are literally chicken tenders.", "It was an emotional wedding. Even the cake was in tiers.", "I just got kicked out of a secret cooking society. I spilled the beans.", "What's a frog's favorite type of shoes? Open toad sandals.", "Blunt pencils are really pointless.", "6:30 is the best time on a clock, hands down.", "Two wifi engineers got married. The reception was fantastic.", "Just got fired from my job as a set designer. I left without making a scene.", "What's the difference between ignorance and apathy? I don't know and I don't care.", "One of the cows didn't produce milk today. It was an udder failure.", "Adam & Eve were the first ones to ignore the Apple terms and conditions.", "Refusing to go to the gym is a form of resistance training.", "If attacked by a mob of clowns, go for the juggler.", "The man who invented Velcro has died. RIP.", "Despite the high cost of living, it remains popular.", "A dung beetle walks into a bar and asks, 'Is this stool taken?'", "I can tell when people are being judgmental just by looking at them.", "The rotation of Earth really makes my day.", "Well, to be Frank with you, I'd have to change my name.", "My friend was explaining electricity to me, but I was like, 'Watt?'", "What if there were no hypothetical questions?", "Are people born with photographic memories, or does it take time to develop?", "The world champion tongue twister got arrested. I hear they're going to give him a tough sentence.", "Pollen is what happens when flowers can't keep it in their plants.", "A book fell on my head the other day. I only have my shelf to blame though.", "Communist jokes aren't funny unless everyone gets them.", "Geology rocks, but geography's where it's at.", "I buy all my guns from a guy called T-Rex. He's a small arms dealer.", "My friend's bakery burned down last night. Now his business is toast.", "Four fonts walk into a bar. The bartender says, 'Hey! We don't want your type in here!'", "If you don't pay your exorcist, do you get repossessed?", "When the cannibal showed up late to the buffet, they gave him the cold shoulder.", "Fighting for peace is like screwing for virginity.", "A ghost walked into a bar and ordered a shot of vodka. The bartender said, 'Sorry, we don't serve spirits here.'", "The man who invented knock-knock jokes should get a no bell prize.", "I bought the world's worst thesaurus yesterday. Not only is it terrible, it's also terrible.", "A blind man walked into a bar… and a table… and a chair…", "A Freudian slip is when you mean one thing and mean your mother.", "I went to a seafood disco last week, but ended up pulling a mussel.", "The first time I got a universal remote control, I thought to myself, 'This changes everything.'", "How do you make holy water? You boil the hell out of it.", "I saw a sign the other day that said, 'Watch for children,' and I thought, 'That sounds like a fair trade.'", "Whiteboards are remarkable.", "I threw a boomerang a couple years ago; I know live in constant fear.", "I put my grandma on speed dial the other day. I call it insta-gram.", "I have a few jokes about unemployed people, but none of them work.", "'I have a split personality,' said Tom, being Frank.", "My teachers told me I'd never amount to much because I procrastinate so much. I told them, “Just you wait!”", "Will glass coffins be a success? Remains to be seen.", "Did you hear about the guy whose whole left side got amputated? He's all right now.", "The man who survived both mustard gas and pepper spray is a seasoned veteran now.", "Have you heard about the new restaurant called 'Karma?' There's no menu—you get what you deserve."}

	choose := rInt(0, len(signnames))

	return signnames[choose]

}
func usescroll() { //MARK: usescroll

	zmagic := xmagic{}
	zmagic.atk = activscroll.atk + player.intel

	switch activscroll.scrollnum {
	case 4: //poison gas
		magictimer = fps * 3

		num := rInt(20, 30)

		for {

			zcircle := xcircle{}
			zcircle.v2 = player.cnt
			zcircle.v2.X += rFloat32(-tilesize*6, tilesize*6)
			zcircle.v2.Y += rFloat32(-tilesize*6, tilesize*6)
			zcircle.rad = rFloat32(tilesize/2, tilesize*2)

			zcircle.dirx = rFloat32(-tilesize/6, tilesize/6)
			zcircle.diry = rFloat32(-tilesize/6, tilesize/6)

			zcircle.color = randomgreen()

			zmagic.circles = append(zmagic.circles, zcircle)

			num--
			if num == 0 {
				break
			}
		}

		activmagic = append(activmagic, zmagic)

		magicon = true
	case 3: //meteors

		magictimer = fps * 3
		num := rInt(3, 6)

		for {

			zmagic := xmagic{}
			zmagic.rec = player.rec

			zmagic.rec.Width = tilesize * 3
			zmagic.rec.Height = tilesize * 3

			zmagic.rec.X += rFloat32(-tilesize*2, tilesize*2)
			zmagic.rec.Y += rFloat32(-tilesize*2, tilesize*2)

			zmagic.dirx = rFloat32(-tilesize/4, tilesize/4)
			zmagic.diry = rFloat32(-tilesize/4, tilesize/4)
			zmagic.ro = rFloat32(0, 360)
			zmagic.atk = activscroll.atk + player.intel
			activmagic = append(activmagic, zmagic)

			num--
			if num == 0 {
				break
			}
		}

		magicon = true
	case 2: //identify
		for a := 0; a < len(inven); a++ {
			inven[a].ident = true
		}
		for a := 0; a < len(beltinven); a++ {
			beltinven[a].ident = true
		}
	case 1: //teleport
		magictimer = fps * 2
		zcircle := xcircle{}

		zcircle.rad = scrhf32 / 2
		zcircle.color = randomorange()
		zmagic.circles = append(zmagic.circles, zcircle)
		zcircle.rad -= 100
		zmagic.circles = append(zmagic.circles, zcircle)
		zcircle.rad -= 100
		zmagic.circles = append(zmagic.circles, zcircle)

		activmagic = append(activmagic, zmagic)
		magicon = true

		player.cnt, _ = findcntr()
		selpoint = player.cnt
		selpoint.X += tilesize / 4

		selrec = rl.NewRectangle(selpoint.X-tilesize/8, selpoint.Y-tilesize/8, tilesize/4, tilesize/4)
		player.dirx = 0
		player.diry = 0
		upplayer()

	case 0: //ring of fire
		magictimer = fps * 2

		zcircle := xcircle{}

		zcircle.rad = 10
		zcircle.color = randomorange()
		zmagic.circles = append(zmagic.circles, zcircle)
		zcircle.rad = 60
		zmagic.circles = append(zmagic.circles, zcircle)
		zcircle.rad = 110
		zmagic.circles = append(zmagic.circles, zcircle)

		activmagic = append(activmagic, zmagic)
		magicon = true

	}

}
func moveroom(roomnum int, x, y float32) { //MARK: moveroom

	//find x y change
	var xchange []float32
	var ychange []float32

	for a := 0; a < len(level[roomnum].roomrec); a++ {
		if level[roomnum].roomrec[a].rec.X == level[roomnum].boundrec.X {
			xchange = append(xchange, 0)
		} else if level[roomnum].roomrec[a].rec.X > level[roomnum].boundrec.X {
			xchange = append(xchange, level[roomnum].roomrec[a].rec.X-level[roomnum].boundrec.X)
		} else if level[roomnum].roomrec[a].rec.X < level[roomnum].boundrec.X {
			xdiff := level[roomnum].boundrec.X - level[roomnum].roomrec[a].rec.X
			xdiff = -xdiff
			xchange = append(xchange, xdiff)
		}

		if level[roomnum].roomrec[a].rec.Y == level[roomnum].boundrec.Y {
			ychange = append(ychange, 0)
		} else if level[roomnum].roomrec[a].rec.Y > level[roomnum].boundrec.Y {
			ychange = append(ychange, level[roomnum].roomrec[a].rec.Y-level[roomnum].boundrec.Y)
		} else if level[roomnum].roomrec[a].rec.Y < level[roomnum].boundrec.Y {
			ydiff := level[roomnum].boundrec.Y - level[roomnum].roomrec[a].rec.Y
			ydiff = -ydiff
			ychange = append(ychange, ydiff)
		}
	}

	//move rooms
	for a := 0; a < len(level[roomnum].roomrec); a++ {

		level[roomnum].roomrec[a].rec.X = x
		level[roomnum].roomrec[a].rec.X += xchange[a]

		level[roomnum].roomrec[a].rec.Y = y
		level[roomnum].roomrec[a].rec.Y += ychange[a]

		level[roomnum].roomrec[a].cnt = makereccnt(level[roomnum].roomrec[a].rec)
		level[roomnum].roomrec[a].collisrec = makecollisrec(level[roomnum].roomrec[a].rec)

	}

	//move boundrec
	level[roomnum].boundrec.X = x
	level[roomnum].boundrec.Y = y

}
func animate() { //MARK: animate

	if animateplayer {
		if frames%8 == 0 {
			playerimg.X += 16
			if playerimg.X > 65 {
				playerimg.X = 1
			}
			playerlimg.X -= 16
			if playerlimg.X < 1 {
				playerlimg.X = 66
			}
		}
	}

}
func checkreccollisions(v2 rl.Vector2) bool { //MARK: checkreccollisions

	canmove := false

	for a := 0; a < len(visroom); a++ {

		for b := 0; b < len(visroom[a].roomrec); b++ {

			if rl.CheckCollisionPointRec(v2, visroom[a].roomrec[b].rec) {
				canmove = true

			}
		}

	}

	if canmove {

		for a := 0; a < len(visroom); a++ {

			for b := 0; b < len(visroom[a].objs); b++ {
				if rl.CheckCollisionPointRec(v2, visroom[a].objs[b].rec) && visroom[a].objs[b].solid {
					canmove = false
				}
			}

		}

	}
	return canmove

}

func getpointdirs(cnt, point rl.Vector2, vel float32) (dirx, diry float32) { //MARK: getpointdirs

	x, y := float32(0), float32(0)

	xdiff := absdiff32(cnt.X, point.X)
	ydiff := absdiff32(cnt.Y, point.Y)

	if xdiff > ydiff {
		x = vel
		if point.X < cnt.X {
			x = -x
		}
		ychange := xdiff / vel
		y = ydiff / ychange
		if point.Y < cnt.Y {
			y = -y
		}
	} else {
		y = vel
		if point.Y < cnt.Y {
			y = -y
		}
		xchange := ydiff / vel
		x = xdiff / xchange
		if point.X < cnt.X {
			x = -x
		}

	}

	return x, y

}
func bounceweapon(bulletnum int) { //MARK: bounceweapon

	if activweapons[bulletnum].dirx < 0 {
		activweapons[bulletnum].dirx = getabs(activweapons[bulletnum].dirx)
	} else {
		activweapons[bulletnum].dirx = -activweapons[bulletnum].dirx
	}

	if activweapons[bulletnum].diry < 0 {
		activweapons[bulletnum].diry = getabs(activweapons[bulletnum].diry)
	} else {
		activweapons[bulletnum].diry = -activweapons[bulletnum].diry
	}

}
func checkarmorset() { //MARK: checkarmorset

	for a := 0; a < len(armorsetcount); a++ {
		armorsetcount[a] = 0
	}

	for a := 0; a < len(inven); a++ {
		if inven[a].legendary && inven[a].invenselect {
			switch inven[a].armorsetnum {
			case 1:
				armorsetcount[0]++
			case 2:
				armorsetcount[1]++
			case 3:
				armorsetcount[2]++
			case 4:
				armorsetcount[3]++
			case 5:
				armorsetcount[4]++
			case 6:
				armorsetcount[5]++
			}
		}
	}

	armorsetactive = false
	for a := 0; a < len(armorsetcount); a++ {

		if armorsetcount[a] > 3 {
			armorsetactive = true
			break
		}

	}

	if armorsetactive {
		uparmorset()
	}

}
func armorsetoff(invennum int) { //MARK: armorsetoff

	switch inven[invennum].armorsetnum {

	case 1:
		player.hpmax -= 50
		if player.hp > player.hpmax {
			player.hp = player.hpmax
		}
		newmsg("armor set of health unequipped... -50 MAX HP")

	case 2:
		player.str -= 10

		player.intel -= 10

		player.dex -= 10

		player.luck -= 10

		newmsg("armor set of boosting unequipped... -10 STR INT DEX LUK")

	case 3:
		player.luck = player.luckorig
		newmsg("lucky luke armor set unequipped... LUK returns to normal")

	case 4:
		player.dex = player.dexorig
		newmsg("robin hood armor set unequipped... DEX returns to normal")

	case 5:
		player.str = player.strorig
		newmsg("the hulk armor set unequipped... STR returns to normal")

	case 6:
		player.intel = player.intelorig
		newmsg("merlin armor set unequipped... INT returns to normal")
	}

	armorsetactive = false

}

// MARK: DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW
func drawhelp() { //MARK: drawhelp

	rl.DrawText("help", 20, 10, 80, rl.Black)

	x := int32(10)
	y := int32(120)

	rl.DrawText("There is an issue/bug and sometimes the level will generate and you will be stuck in a room. Press the STUCK button bottom right and the next level will generate.", x, y, txtdef, rl.Black)
	y += txtdef + txtdef/2

	rl.DrawText("The entire game can be played using only a mouse, there is no option to use a controller or keyboard.", x, y, txtdef, rl.Black)
	y += txtdef * 3

	rl.DrawText("Left Mouse Click : ", x+40, y, txtdef, rl.Black)
	rl.DrawText("On open floor moves to that area", x+250, y, txtdef, rl.Black)
	y += txtdef + txtdef/2
	rl.DrawText("On an interactable object will pick up or use the object", x+250, y, txtdef, rl.Black)
	y += txtdef + txtdef/2
	rl.DrawText("On inventory will equip/unequip or use an item", x+250, y, txtdef, rl.Black)

	y += txtdef * 3

	rl.DrawText("Right Mouse Click : ", x+40, y, txtdef, rl.Black)
	rl.DrawText("With weapon equipped will attack", x+250, y, txtdef, rl.Black)
	y += txtdef + txtdef/2
	rl.DrawText("With spade equipped will dig (only works in certain areas)", x+250, y, txtdef, rl.Black)
	y += txtdef + txtdef/2
	rl.DrawText("On inventory will open options to delete or move the item", x+250, y, txtdef, rl.Black)

	y += txtdef * 3

	rl.DrawText("Keys : ", x+40, y, txtdef, rl.Black)
	rl.DrawText("ESC = Exits the game", x+250, y, txtdef, rl.Black)
	y += txtdef + txtdef/2
	rl.DrawText("PAUSE = Does what it says", x+250, y, txtdef, rl.Black)
	y += txtdef + txtdef/2
	rl.DrawText("M = Open/Closes Map (to teleport on the map click in a blue rectangle)", x+250, y, txtdef, rl.Black)
	y += txtdef + txtdef/2
	rl.DrawText("L = Open/Closes Message Log", x+250, y, txtdef, rl.Black)
	y += txtdef + txtdef/2
	rl.DrawText("H = Open/Closes Help (this screen)", x+250, y, txtdef, rl.Black)

	y += txtdef * 3

	rl.DrawText("Most objects in the word have some purpose and can be interacted with so explore and interact to work out how to stay alive.", x, y, txtdef, rl.Black)
	y += txtdef + txtdef/2

	rl.DrawText("If you have gems in your inventory at the end of the level (when going downstairs), if the combined gem value is high enough the shop will open.", x, y, txtdef, rl.Black)
	y += txtdef + txtdef/2

	rl.DrawText("Belt inventory only allows small items - potions, maps & scrolls / Backpack inventory allows all items.", x, y, txtdef, rl.Black)
	y += txtdef + txtdef/2

	rl.DrawText("The KILLS TO BOSS (bottom left) text is a countdown of kills till the next boss will appear onscreen, when that number reaches zero a random boss will appear.", x, y, txtdef, rl.Black)
	y += txtdef + txtdef/2

	rl.DrawText("Each level has more monsters than the previous meaning the further you go the more difficult the game becomes.", x, y, txtdef, rl.Black)
	y += txtdef + txtdef/2

	rl.DrawText("Pets do nothing other than keep you company so don't bother trying to work out what they do.", x, y, txtdef, rl.Black)
	y += txtdef + txtdef/2

	rl.DrawText("There is no ending, the name of the game is THE ENDLESS DUNGEONS OF PIXEL hence it is endless.", x, y, txtdef, rl.Black)
	y += txtdef + txtdef/2

	//close
	rec := rl.NewRectangle(0, 0, (tilesize + tilesize + tilesize/3), (tilesize/3)*2)

	rec.X = scrwf32 - (rec.Width + tilesize)
	rec.Y = tilesize / 3

	txt := "close"
	txtlen := rl.MeasureText(txt, txtdef)

	if rl.CheckCollisionPointRec(mousev2, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.White, fadeblink))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			helpon = false
			pause = false
		}
	}
	rl.DrawRectangleLinesEx(rec, 4, rl.Black)
	rl.DrawText(txt, rec.ToInt32().X+rec.ToInt32().Width/2-txtlen/2, rec.ToInt32().Y+5, txtdef, rl.Black)

}
func drawendgame() { //MARK: drawendgame

	rl.DrawRectangle(0, 0, scrw, scrh, rl.Fade(rl.Black, 0.7))

	closerec := rl.NewRectangle(scrwf32/2-100, scrhf32/2-50, 200, 100)

	rl.SetExitKey(rl.KeyY)

	rl.DrawRectangleRec(closerec, brightorange())
	rl.DrawText("press key to exit", closerec.ToInt32().X+10, closerec.ToInt32().Y+5, 20, rl.Black)
	rl.DrawText("Y / N", closerec.ToInt32().X+20, closerec.ToInt32().Y+40, 50, rl.Black)

	if rl.IsKeyPressed(rl.KeyN) {
		pause = false
		endgamewindow = false
	}

}
func drawshop() { //MARK: drawshop

	rl.DrawText("shop", 20, 10, 80, rl.Black)
	destrec := rl.NewRectangle(scrwf32-300, 10, tilesize*2, tilesize*2)

	//key gem imgs
	rl.DrawTexturePro(imgs, shopgemimg, destrec, origin, 0, randomcolor())
	rl.DrawText(fmt.Sprint(gemstotal), destrec.ToInt32().X+destrec.ToInt32().Width+int32(tilesize/2), destrec.ToInt32().Y+int32(tilesize/4), txtxl, rl.Black)
	destrec.X -= tilesize * 5
	rl.DrawTexturePro(imgs, shopkeyimg, destrec, origin, 0, randomcolor())
	rl.DrawText(fmt.Sprint(keystotal), destrec.ToInt32().X+destrec.ToInt32().Width+int32(tilesize/2), destrec.ToInt32().Y+int32(tilesize/4), txtxl, rl.Black)

	//shop recs
	for a := 0; a < len(shoprecs); a++ {
		rl.DrawRectangleLinesEx(shoprecs[a], 16, rl.Black)
		rl.DrawRectangleRec(shoprecs[a], rl.Fade(rl.Black, 0.4))

		destrec := rl.NewRectangle(shoprecs[a].X, shoprecs[a].Y, tilesize*5, tilesize*5)

		//keys locked
		if shopitems[a].locked {

			rl.DrawTexturePro(imgs, lockimg, destrec, origin, 0, randombluelight())

			if rl.CheckCollisionPointRec(mousev2, shoprecs[a]) {
				txtlen := rl.MeasureText("use a key to unlock this item", 40)

				rl.DrawText("use a key to unlock this item", scrw/2-txtlen/2, int32(shoprecs[a].Y-60), 40, rl.Black)

				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
					if keystotal > 0 {
						keystotal--
						shopitems[a].locked = false

						found := false
						for a := 0; a < len(inven); a++ {

							if inven[a].name == "key" {
								inven[a] = xobj{}
								findinvennum()
								found = true
								break
							}
						}
						if !found {
							for a := 0; a < len(beltinven); a++ {

								if beltinven[a].name == "key" {
									beltinven[a] = xobj{}
									findbeltinvennum()
									break
								}
							}

						}

					} else {
						nokeys = true
						nokeystimer = fps * 3
					}
				}

			}

		} else { //unlocked
			if shopitems[a].sold {
				rl.DrawTexturePro(imgs, soldimg, destrec, origin, 0, rl.White)
			} else {

				switch shopitems[a].name2 {
				case "coins":
					rl.DrawTexturePro(imgs, coinimg, destrec, origin, 0, rl.White)
				case "cactus":
					rl.DrawTexturePro(imgs, shopitems[a].img, destrec, origin, 0, randomgreen())
				case "disease":
					rl.DrawTexturePro(imgs, shopitems[a].img, destrec, origin, 0, randomyellow())
				case "fireresist":
					rl.DrawTexturePro(imgs, shopitems[a].img, destrec, origin, 0, randomorange())
				case "firetrail":
					rl.DrawTexturePro(imgs, shopitems[a].img, destrec, origin, 0, randomorange())
				case "poison":
					rl.DrawTexturePro(imgs, shopitems[a].img, destrec, origin, 0, randomgreen())
				case "tel99":
					origin := rl.NewVector2(destrec.Width/2, destrec.Height/2)
					destrec.X += destrec.Width / 2
					destrec.Y += destrec.Height / 2
					rl.DrawTexturePro(imgs, shopitems[a].img, destrec, origin, shopitems[a].ro, rl.White)
					shopitems[a].ro += 2
				default:
					rl.DrawTexturePro(imgs, shopitems[a].img, destrec, origin, 0, shopitems[a].color)
				}

				if rl.CheckCollisionPointRec(mousev2, shoprecs[a]) {
					txtlen := rl.MeasureText(shopitems[a].name+"  gem value: "+fmt.Sprint(shopitems[a].cost), 40)

					rl.DrawText(shopitems[a].name+"  gem value: "+fmt.Sprint(shopitems[a].cost), scrw/2-txtlen/2, int32(shoprecs[a].Y-60), 40, rl.Black)

					if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
						if gemstotal >= shopitems[a].cost {
							gemstotal -= shopitems[a].cost
							applyshopitem(a)
							shopitems[a].sold = true
						}
					}

				}
			}
		}

		if nokeys {
			if nokeystimer > 0 {
				txtlen := rl.MeasureText("you have no keys in your inventory", 40)
				rl.DrawText("you have no keys in your inventory", scrw/2-txtlen/2, int32(shoprecs[a].Y+shoprecs[a].Height+20), 40, rl.Black)
			}

		}

		if !closeshop {
			if shoprecs[0].X > (scrwf32/2)-(shoprectotallen/2) {
				shoprecs[a].X -= tilesize / 2
			}
		}
	}

	//close
	x := scrwf32 - 200
	y := scrhf32 - 100

	rec := rl.NewRectangle(x, y, (tilesize + tilesize + tilesize/3), (tilesize/3)*2)

	txt := "close"
	txtlen := rl.MeasureText(txt, txtdef)

	if rl.CheckCollisionPointRec(mousev2, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.White, fadeblink))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			closeshop = true
		}
	}
	rl.DrawRectangleLinesEx(rec, 4, rl.Black)
	rl.DrawText(txt, rec.ToInt32().X+rec.ToInt32().Width/2-txtlen/2, int32(y)+4, txtdef, rl.Black)

	if closeshop {
		for a := 0; a < len(shoprecs); a++ {
			if shoprecs[4].X+shoprecs[4].Width > -tilesize {
				shoprecs[a].X -= tilesize / 2
			} else {
				upgemvalues()
				if randomarmoron {

					zobj := xobj{}
					zobj.cnt = player.cnt
					zobj.cnt.X += tilesize * 2

					zobj.kind = "armor"
					zobj.collect = true
					zobj.questitem = true
					zobj.legendary = true
					zobj.name = "legendary armor"
					zobj.color = brightyellow()
					zobj.armorsetnum = rInt(1, 7)
					zobj.ability = rInt(1, 9)

					switch zobj.ability {
					case 1:
						zobj.name4 = "health regen"
					case 2:
						zobj.name4 = "vampirism"
					case 3:
						zobj.name4 = "thorns"
					case 4:
						zobj.name4 = "speed"
					case 5:
						zobj.name4 = "fire trail"
					case 6:
						zobj.name4 = "teleport"
						zobj.timer = rInt32(int(fps*10), int(fps*20))
					case 7:
						zobj.name4 = "rainbow"
					case 8:
						zobj.name4 = "identify"

					}

					choose := rInt(1, 7)
					switch choose {
					case 1:
						zobj.img = helmetimgs[rInt(0, len(helmetimgs))]
						zobj.usetype = 1

						zobj.name2 = "legendary helmet of " + zobj.name4 + " - part of armor set " + (fmt.Sprint(zobj.armorsetnum))
						zobj.name3 = "helmet"
					case 2:
						zobj.img = bootimgs[rInt(0, len(bootimgs))]
						zobj.usetype = 2

						zobj.name2 = "legendary boots of " + zobj.name4 + " - part of armor set " + (fmt.Sprint(zobj.armorsetnum))
						zobj.name3 = "boots"
					case 3:
						zobj.img = gloveimgs[rInt(0, len(gloveimgs))]
						zobj.usetype = 3

						zobj.name2 = "legendary gloves of " + zobj.name4 + " - part of armor set " + (fmt.Sprint(zobj.armorsetnum))
						zobj.name3 = "gloves"
					case 4:
						zobj.img = vestimgs[rInt(0, len(vestimgs))]
						zobj.usetype = 4

						zobj.name2 = "legendary vest of " + zobj.name4 + " - part of armor set " + (fmt.Sprint(zobj.armorsetnum))
						zobj.name3 = "vest"
					case 5:
						zobj.img = robeimgs[rInt(0, len(robeimgs))]
						zobj.usetype = 4

						zobj.name2 = "legendary robe of " + zobj.name4 + " - part of armor set " + (fmt.Sprint(zobj.armorsetnum))
						zobj.name3 = "vest"
					case 6:
						zobj.img = crownimgs[rInt(0, len(crownimgs))]
						zobj.usetype = 1

						zobj.name2 = "legendary crown of " + zobj.name4 + " - part of armor set " + (fmt.Sprint(zobj.armorsetnum))
						zobj.name3 = "helmet"
					}

					zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
					zobj.boundrec = zobj.rec
					zobj.boundrec.X -= tilesize * 2
					zobj.boundrec.Y -= tilesize * 2
					zobj.boundrec.Width += tilesize * 4
					zobj.boundrec.Height += tilesize * 4
					level[0].objs = append(level[0].objs, zobj)

					randomarmoron = false
				}
				pause = false
				shopon = false
				closeshop = false
			}
		}

	}

}

func drawsettings() { //MARK: drawsettings

	rl.DrawText("settings", 20, 10, 80, rl.Black)

	x := float32(40)
	y := float32(150)

	txt := "crt scan lines"
	rl.DrawText(txt, int32(x), int32(y), txtl, rl.Black)
	x += tilesize * 11
	scan = drawtoggle(x, y+5, scan)
	x = 40
	y += tilesize

	txt = "sprite ghosting"
	rl.DrawText(txt, int32(x), int32(y), txtl, rl.Black)
	x += tilesize * 11
	ghost = drawtoggle(x, y+5, ghost)
	x = 40
	y += tilesize

	txt = "auto switch weapons"
	rl.DrawText(txt, int32(x), int32(y), txtl, rl.Black)
	x += tilesize * 11
	autoswitchweapons = drawtoggle(x, y+5, autoswitchweapons)
	x = 40
	y += tilesize

	txt = "auto switch ammo"
	rl.DrawText(txt, int32(x), int32(y), txtl, rl.Black)
	x += tilesize * 11
	autoswitchammo = drawtoggle(x, y+5, autoswitchammo)
	x = 40
	y += tilesize

	txt = "map teleport"
	rl.DrawText(txt, int32(x), int32(y), txtl, rl.Black)
	x += tilesize * 11
	teleporton = drawtoggle(x, y+5, teleporton)
	x = 40
	y += tilesize

	txt = "mr pigeon"
	rl.DrawText(txt, int32(x), int32(y), txtl, rl.Black)
	x += tilesize * 11
	pigeononoff = drawtoggle(x, y+5, pigeononoff)
	x = 40
	y += tilesize

	txt = "show enemy hp"
	rl.DrawText(txt, int32(x), int32(y), txtl, rl.Black)
	x += tilesize * 11
	showmonshp = drawtoggle(x, y+5, showmonshp)
	x = 40
	y += tilesize

	txt = "invincible mode"
	rl.DrawText(txt, int32(x), int32(y), txtl, rl.Black)
	x += tilesize * 11
	switchinvincible = drawtoggle(x, y+5, switchinvincible)
	x = 40
	y += tilesize

	txt = "mute all audio"
	rl.DrawText(txt, int32(x), int32(y), txtl, rl.Black)
	x += tilesize * 11
	mute = drawtoggle(x, y+5, mute)
	x = 40
	y += tilesize

	txt = "sound fx"
	rl.DrawText(txt, int32(x), int32(y), txtl, rl.Black)
	x += tilesize * 11
	soundfxon = drawtoggle(x, y+5, soundfxon)
	x = 40
	y += tilesize

	txt = "music"
	rl.DrawText(txt, int32(x), int32(y), txtl, rl.Black)
	x += tilesize * 11
	musicon = drawtoggle(x, y+5, musicon)
	x = 40
	y += tilesize

	txt = "volume"
	rl.DrawText(txt, int32(x), int32(y), txtl, rl.Black)
	x += tilesize*9 + (tilesize / 4)

	rec := rl.NewRectangle(x, y, tilesize/2, tilesize/2)

	rl.DrawRectangleRec(rec, rl.Black)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) && clickpause == 0 {

			clickpause = fps / 4

			if soundvol == 1.0 {
				soundvol = 0.9
			} else if soundvol == 0.9 {
				soundvol = 0.8
			} else if soundvol == 0.8 {
				soundvol = 0.7
			} else if soundvol == 0.7 {
				soundvol = 0.6
			} else if soundvol == 0.6 {
				soundvol = 0.5
			} else if soundvol == 0.5 {
				soundvol = 0.4
			} else if soundvol == 0.4 {
				soundvol = 0.3
			} else if soundvol == 0.3 {
				soundvol = 0.2
			} else if soundvol == 0.2 {
				soundvol = 0.1
			} else if soundvol == 0.1 {
				soundvol = 0.0
			}

			rl.SetMasterVolume(soundvol)
		}
	}
	rl.DrawText("-", rec.ToInt32().X+int32(tilesize/5), rec.ToInt32().Y+int32(tilesize/9), txtdef, rl.White)

	x += (tilesize / 4) * 3

	txt = fmt.Sprintf("%.0f", soundvol*100)
	rl.DrawText(txt, int32(x), int32(y), txtl, rl.Black)

	x += tilesize

	rec.X = x

	rl.DrawRectangleRec(rec, rl.Black)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) && clickpause == 0 {

			clickpause = fps / 4

			if soundvol == 0.0 {
				soundvol = 0.1
			} else if soundvol == 0.1 {
				soundvol = 0.2
			} else if soundvol == 0.2 {
				soundvol = 0.3
			} else if soundvol == 0.3 {
				soundvol = 0.4
			} else if soundvol == 0.4 {
				soundvol = 0.5
			} else if soundvol == 0.5 {
				soundvol = 0.6
			} else if soundvol == 0.6 {
				soundvol = 0.7
			} else if soundvol == 0.7 {
				soundvol = 0.8
			} else if soundvol == 0.8 {
				soundvol = 0.9
			} else if soundvol == 0.9 {
				soundvol = 1.0
			}

			rl.SetMasterVolume(soundvol)
		}
	}
	rl.DrawText("+", rec.ToInt32().X+int32(tilesize/5), rec.ToInt32().Y+int32(tilesize/9), txtdef, rl.White)

	//close
	rec = rl.NewRectangle(x, y, (tilesize + tilesize + tilesize/3), (tilesize/3)*2)

	x = scrwf32 - (rec.Width + tilesize)
	y = tilesize / 3
	rec.X = x
	rec.Y = y

	txt = "close"
	txtlen := rl.MeasureText(txt, txtdef)

	if rl.CheckCollisionPointRec(mousev2, rec) {
		rl.DrawRectangleRec(rec, rl.Fade(rl.White, fadeblink))
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			settingson = false
			pause = false
		}
	}
	rl.DrawRectangleLinesEx(rec, 4, rl.Black)
	rl.DrawText(txt, rec.ToInt32().X+rec.ToInt32().Width/2-txtlen/2, int32(y)+5, txtdef, rl.Black)

}
func drawdied() { //MARK: drawdied

	settingson = false
	mapon = false

	rec := rl.NewRectangle(0, diedy, scrwf32, scrhf32)

	rl.DrawRectangleRec(rec, rl.Black)

	if diedy < 0 {
		diedy += 10
	} else {
		txtlen := rl.MeasureText("you died", 80)
		rl.DrawText("you died", scrw/2-(txtlen/2), scrh/2-40, 80, rl.Fade(rl.White, diedfade))

		if rolldice()+rolldice() == 12 {
			if flipcoin() {
				rl.DrawText("you died", scrw/2-(txtlen/2), scrh/2-40, 80, rl.Fade(brightred(), diedfade))
			} else {
				rl.DrawText("you died", scrw/2-(txtlen/2), scrh/2-40, 80, rl.Fade(brightorange(), diedfade))
			}
		}

		if diedfade < 1.0 {
			diedfade += 0.05
		} else {
			txtlen := rl.MeasureText("click left mouse button", 20)
			rl.DrawText("click left mouse button", scrw/2-(txtlen/2), scrh/2+80, 20, rl.White)

			if rolldice()+rolldice() == 12 {
				rl.DrawText("click left mouse button", scrw/2-(txtlen/2), scrh/2+80, 20, randomcolor())
			}

			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				scoreon = true
				scoreontimer = 0
			}

		}
	}

	//close game
	closerec := rl.NewRectangle(scrwf32-(tilesize/2+tilesize/8), tilesize/8, tilesize/2, tilesize/2)
	rl.DrawRectangleRec(closerec, brightorange())
	rl.DrawLine(closerec.ToInt32().X, closerec.ToInt32().Y, closerec.ToInt32().X+int32(tilesize/2), closerec.ToInt32().Y+int32(tilesize/2), rl.Black)
	rl.DrawLine(closerec.ToInt32().X+int32(tilesize/2), closerec.ToInt32().Y, closerec.ToInt32().X, closerec.ToInt32().Y+int32(tilesize/2), rl.Black)
	if rl.CheckCollisionPointRec(mousev2, closerec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			pause = true
			endgamewindow = true
		}
	}

	//scanlines

	for a := 0; a < len(scanlines); a++ {
		rl.DrawLineV(scanlines[a].v1, scanlines[a].v2, rl.Fade(rl.Black, 0.7))
		scanlines[a].v1.Y += 1
		scanlines[a].v2.Y = scanlines[a].v1.Y

		if scanlines[a].v1.Y == scrhf32 {
			scanlines[a].v1.Y = 0
			scanlines[a].v2.Y = 0
		}
	}

}
func drawintro() { //MARK: drawintro

	if introtxtx < scrw && !intro3on {
		introtxtx += 10
		rl.DrawText("made with", introtxtx, scrh/2-30, 60, rl.White)
		if rl.IsKeyPressed(rl.KeySpace) {
			introtxtx = scrw + 2
		}
	} else {

		destrec := rl.NewRectangle(0, 0, (gologo.Width / 2), (gologo.Height / 2))

		destrec.X += scrwf32 / 2
		destrec.X -= destrec.Width / 2

		destrec.Y += scrhf32 / 2
		destrec.Y -= destrec.Height / 2

		rl.DrawTexturePro(imgs, gologo, destrec, origin, 0, rl.Fade(rl.White, introfade))
		if introfade < 1 && !intro2on {
			introfade += 0.01
		} else {
			intro2on = true
		}

	}

	destrec := rl.NewRectangle(0, 0, (rayliblogo.Width), (rayliblogo.Height))

	if intro2on && !intro3on {

		if introfade > 0 {
			introfade -= 0.01
		} else {

			destrec.X += scrwf32 / 2
			destrec.X -= destrec.Width / 2

			destrec.Y += scrhf32 / 2
			destrec.Y -= destrec.Height / 2

			rl.DrawTexturePro(imgs, rayliblogo, destrec, origin, 0, rl.Fade(rl.White, introfade2))

			if introfade2 < 1 && !intro3on {
				introfade2 += 0.01
			} else {
				intro3on = true
			}

		}

	}

	if intro3on && !intro4on {

		if introfade2 > 1 {
			introfade2 -= 0.01
			rl.DrawTexturePro(imgs, rayliblogo, destrec, origin, 0, rl.Fade(rl.White, introfade2))
		} else {

			if introtxtx2 < scrw {
				introtxtx2 += 10
				rl.DrawText("a game by nicholasimon", introtxtx2, scrh/2-30, 60, rl.White)
				if rl.IsKeyPressed(rl.KeySpace) {
					introtxtx2 = scrw + 2
				}
			} else {
				intro4on = true
			}
		}
	}

	if intro4on {
		rl.PlayMusicStream(intromusic)

		rl.UpdateMusicStream(intromusic)

		rec := rl.NewRectangle(0, 0, scrwf32, scrhf32)

		rl.DrawRectangleRec(rec, rl.Black)

		rl.DrawText("the", 100, introtxty+100, 40, rl.White)
		rl.DrawText("endless dungeons", 100, introtxty+150, 120, rl.White)
		rl.DrawText("of", 100, introtxty+290, 40, rl.White)
		rl.DrawText("pixel", 100, introtxty+340, 120, rl.White)

		if introtxty < 0 {
			introtxty += 10
		} else {

			if frames%8 == 0 {
				playerimg.X += 16
				if playerimg.X > 65 {
					playerimg.X = 1
				}
				playerlimg.X -= 16
				if playerlimg.X < 1 {
					playerlimg.X = 66
				}
			}

			destrec := rl.NewRectangle(introplayx, scrhf32-tilesize*4, tilesize*4, tilesize*4)

			if introplayx > -tilesize*5 && !introlr {
				rl.DrawTexturePro(imgs, playerlimg, destrec, origin, 0, player.color)
				introplayx -= 10
			} else if introplayx <= -tilesize*5 && !introlr {
				introtimer = rInt(10, 20)
				introlr = true
			}

			if introlr && introtimer == 0 && introplayx < scrwf32+tilesize {
				rl.DrawTexturePro(imgs, playerimg, destrec, origin, 0, player.color)
				introplayx += 10
			} else if introplayx > scrwf32+tilesize {
				introtimer = rInt(10, 20)
				introlr = false
			}

			txtlen := rl.MeasureText("press space / click left mouse", 40)
			rl.DrawText("press space / click left mouse", scrw/2-txtlen/2, scrh/2+100, 40, rl.White)
			if rolldice() == 6 {
				rl.DrawText("press space / click left mouse", scrw/2-txtlen/2, scrh/2+100, 40, brightorange())
			}

			if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsKeyPressed(rl.KeySpace) {
				introon = false
				pause = false
			}
		}

	}

	if rl.IsKeyPressed(rl.KeySpace) || rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		intro4on = true
	}

	//scanlines

	for a := 0; a < len(scanlines); a++ {
		rl.DrawLineV(scanlines[a].v1, scanlines[a].v2, rl.Fade(rl.Black, 0.7))
		scanlines[a].v1.Y += 1
		scanlines[a].v2.Y = scanlines[a].v1.Y

		if scanlines[a].v1.Y == scrhf32 {
			scanlines[a].v1.Y = 0
			scanlines[a].v2.Y = 0
		}
	}

}

func drawmap() { //MARK: drawmap
	rl.BeginMode2D(cammap)

	//rooms layer 1
	for a := 0; a < len(level); a++ {

		for b := 0; b < len(level[a].roomrec); b++ {

			rl.DrawRectangleRec(level[a].roomrec[b].rec, rl.Black)
			if level[a].roomrec[b].visited {
				rl.DrawRectangleRec(level[a].roomrec[b].rec, rl.Fade(rl.DarkBlue, 0.3))

				if teleporton && player.teleports > 0 {
					mousecam := rl.GetScreenToWorld2D(mousev2, cammap)

					if rl.CheckCollisionPointRec(mousecam, level[a].roomrec[b].rec) {

						rl.DrawRectangleRec(level[a].roomrec[b].rec, rl.Fade(rl.DarkBlue, 0.7))

						rl.DrawCircleV(mousecam, tilesize, rl.White)

						if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

							player.cnt = level[a].roomrec[b].cnt
							selpoint = player.cnt
							selrec = rl.NewRectangle(selpoint.X-tilesize/8, selpoint.Y-tilesize/8, tilesize/4, tilesize/4)
							upplayer()

							activscroll.scrollnum = 1
							zmagic := xmagic{}
							magictimer = fps * 2
							zcircle := xcircle{}

							zcircle.rad = scrhf32 / 2
							zcircle.color = randomorange()
							zmagic.circles = append(zmagic.circles, zcircle)
							zcircle.rad -= 100
							zmagic.circles = append(zmagic.circles, zcircle)
							zcircle.rad -= 100
							zmagic.circles = append(zmagic.circles, zcircle)

							activmagic = append(activmagic, zmagic)
							magicon = true

							player.teleports--
							newmsg("you have teleported...")
							mapon = false
							pause = false

						}
					}
				}
			}

		}

		for b := 0; b < len(level[a].objs); b++ {

			if level[a].objs[b].name == "staircase" {
				rl.DrawCircleV(level[a].objs[b].cnt, 50, rl.Fade(brightred(), fadeblink))
			}

		}

	}
	//other layer 2
	for a := 0; a < len(level); a++ {

		for b := 0; b < len(level[a].objs); b++ {

			if level[a].objs[b].name == "staircase" {
				rl.DrawCircleV(level[a].objs[b].cnt, 50, rl.Fade(brightred(), fadeblink))
			}

		}

	}

	rl.DrawCircleV(player.cnt, 50, rl.Fade(rl.Green, fadeblink))

	rl.EndMode2D()

}
func drawtest() { //MARK: drawtest

	rl.DrawRectangle(0, 0, scrw, scrh, rl.Black)

	rl.DrawCircleV(testv1, tilesize/4, rl.White)
	rl.DrawCircleV(testv2, tilesize/4, rl.White)

	//	angle := angle2points(testv1, testv2)

	testangle = testangle * (math.Pi / 180)

	newx := float32(math.Cos(float64(testangle)))*(testv2.X-testv1.X) - float32(math.Sin(float64(testangle)))*(testv2.Y-testv1.Y) + testv1.X

	newy := float32(math.Sin(float64(testangle)))*(testv2.X-testv1.X) + float32(math.Cos(float64(testangle)))*(testv2.Y-testv1.Y) + testv1.Y

	testv2 = rl.NewVector2(newx, newy)
	testangle++

}
func drawpigeon() { //MARK: drawpigeon

	if pigeonflyingon {

		destrec := rl.NewRectangle(pigeonx, pigeony, tilesize*2, tilesize*2)
		rl.DrawTexturePro(imgs, pigeonflyingimg, destrec, origin, 0, rl.White)
		pigeonx--
		if frames%6 == 0 {
			pigeonflyingimg.X -= pigeonflyingimg.Width
		}
		if pigeonflyingimg.X < 9 {
			pigeonflyingimg.X = 57
		}

		if pigeonx < 0 {
			pigeonflyingon = false
			pigeonon = false
			pigeontimer = fps * rInt32(20, 100)
		}

	} else {

		pigeony = footerrec.Y - tilesize*2
		destrec := rl.NewRectangle(pigeonx, pigeony, tilesize*2, tilesize*2)

		rl.DrawTexturePro(imgs, pigeonimg, destrec, origin, 0, rl.White)

		if rl.CheckCollisionPointRec(mousev2, destrec) {

			txtlen := rl.MeasureText(pigeonphrase, 20)

			rec := rl.NewRectangle(destrec.X+destrec.Width/2-float32(txtlen/2)-2, destrec.Y-24, float32(txtlen+4), 24)

			rl.DrawRectangleRec(rec, rl.Black)

			rl.DrawText(pigeonphrase, destrec.ToInt32().X+destrec.ToInt32().Width/2-txtlen/2, destrec.ToInt32().Y-24, 20, rl.White)

		}

		pigeonx++
		if pigeonx > scrwf32 {
			pigeonflyingon = true
			pigeony = rFloat32(tilesize, tilesize*10)
			pigeonon = false
			pigeontimer = fps * rInt32(20, 100)
		}

		if frames%6 == 0 {
			pigeonimg.X += pigeonimg.Width
		}

		if pigeonimg.X > 56 {
			pigeonimg.X = 8
		}
	}

}

func drawtoggle(x, y float32, name bool) bool { //MARK: drawtoggle

	rec := rl.NewRectangle(x, y, tilesize/2, tilesize/2)

	if name {
		rl.DrawRectangleRec(rec, brightred())
	}
	rl.DrawRectangleLinesEx(rec, 4, rl.Black)

	if rl.CheckCollisionPointRec(mousev2, rec) {

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if name {
				name = false
			} else {
				name = true
			}
		}
	}

	return name

}
func drawblood() { //MARK: drawblood

	clear := true
	for a := 0; a < len(blood); a++ {
		if !blood[a].inactiv {
			clear = false
			rl.DrawCircleV(blood[a].v2, blood[a].rad, rl.Fade(blood[a].color, blood[a].fade))
		}
		blood[a].fade -= rFloat32(0.01, 0.1)

		if blood[a].fade <= 0 {
			blood[a].inactiv = true
		}
	}

	if clear {
		blood = nil
	}

}
func drawmagic() { //MARK: drawmagic

	for a := 0; a < len(activmagic); a++ {

		switch activscroll.scrollnum {

		case 4: //poison gas
			if len(activmagic[a].circles) > 0 {

				for b := 0; b < len(activmagic[a].circles); b++ {

					rl.DrawCircleV(activmagic[a].circles[b].v2, activmagic[a].circles[b].rad, rl.Fade(activmagic[a].circles[b].color, rFloat32(0.3, 0.7)))

					if ghost {
						v3 := activmagic[a].circles[b].v2
						v3.X += rFloat32(-12, 13)
						v3.Y += rFloat32(-12, 13)
						rl.DrawCircleV(v3, activmagic[a].circles[b].rad, rl.Fade(activmagic[a].circles[b].color, rFloat32(0.2, 0.5)))

					}

					activmagic[a].circles[b].v2.X += activmagic[a].circles[b].dirx
					activmagic[a].circles[b].v2.Y += activmagic[a].circles[b].diry

					for c := 0; c < len(vismonsters); c++ {
						if rl.CheckCollisionPointCircle(vismonsters[c].cnt, activmagic[a].circles[b].v2, activmagic[a].circles[b].rad) {

							if !vismonsters[c].poisoned {
								monsters[vismonsters[c].num].poisoned = true
								monsters[vismonsters[c].num].timer = int32(activmagic[a].atk) * fps
								vismonsters[c].poisoned = true
							}

						}
					}

				}

			}
		case 3: //meteors

			origin2 := rl.NewVector2(activmagic[a].rec.Width/2, activmagic[a].rec.Height/2)
			destrec := activmagic[a].rec
			destrec.X += activmagic[a].rec.Width / 2
			destrec.Y += activmagic[a].rec.Height / 2

			rl.DrawTexturePro(imgs, meteorimg, destrec, origin2, activmagic[a].ro, randomgrey())

			if ghost {

				destrec2 := destrec
				destrec2.X += rFloat32(-7, 8)
				destrec2.Y += rFloat32(-7, 8)
				rl.DrawTexturePro(imgs, meteorimg, destrec2, origin2, activmagic[a].ro, rl.Fade(randomgrey(), rFloat32(0.2, 0.7)))

			}

			for c := 0; c < len(vismonsters); c++ {
				if rl.CheckCollisionRecs(vismonsters[c].rec, activmagic[a].rec) {

					if vismonsters[c].hppause == 0 {
						monsters[vismonsters[c].num].hppause = fps / 2
						vismonsters[c].hppause = fps / 2
						monsters[vismonsters[c].num].hp -= activmagic[a].atk
						if monsters[vismonsters[c].num].hp <= 0 {
							monsters[vismonsters[c].num].inactiv = true
							makededmonster(monsters[vismonsters[c].num].cnt)
						}
					}

				}
			}

			//	rl.DrawRectangleLinesEx(activmagic[a].rec, 4, rl.Green)

			activmagic[a].rec.X += activmagic[a].dirx
			activmagic[a].rec.Y += activmagic[a].diry
			activmagic[a].ro += 4

		case 1: //teleport
			if len(activmagic[a].circles) > 0 {

				for b := 0; b < len(activmagic[a].circles); b++ {
					//	rl.DrawCircleLines(int32(player.cnt.X), int32(player.cnt.Y), activmagic[a].circles[b].rad, activmagic[a].circles[b].color)

					rl.DrawCircleLines(int32(player.cnt.X), int32(player.cnt.Y), activmagic[a].circles[b].rad, rl.White)

					activmagic[a].circles[b].rad -= 10

					for c := 0; c < len(vismonsters); c++ {
						if rl.CheckCollisionPointCircle(vismonsters[c].cnt, player.cnt, activmagic[a].circles[b].rad) {

							if vismonsters[c].hppause == 0 {
								monsters[vismonsters[c].num].hppause = fps / 2
								vismonsters[c].hppause = fps / 2
								monsters[vismonsters[c].num].hp -= activmagic[a].atk
								if monsters[vismonsters[c].num].hp <= 0 {
									monsters[vismonsters[c].num].inactiv = true
									makededmonster(monsters[vismonsters[c].num].cnt)
								}
							}

						}
					}

				}
			}

		case 0: //ring of fire
			if len(activmagic[a].circles) > 0 {

				for b := 0; b < len(activmagic[a].circles); b++ {
					//	rl.DrawCircleLines(int32(player.cnt.X), int32(player.cnt.Y), activmagic[a].circles[b].rad, activmagic[a].circles[b].color)

					rl.DrawRing(player.cnt, activmagic[a].circles[b].rad, activmagic[a].circles[b].rad+rFloat32(18, 25), 0, 360, 40, rl.Fade(randomorange(), rFloat32(0.3, 0.7)))

					activmagic[a].circles[b].rad += 5

					for c := 0; c < len(vismonsters); c++ {
						if rl.CheckCollisionPointCircle(vismonsters[c].cnt, player.cnt, activmagic[a].circles[b].rad) {

							if vismonsters[c].hppause == 0 {
								monsters[vismonsters[c].num].hppause = fps / 2
								vismonsters[c].hppause = fps / 2
								monsters[vismonsters[c].num].hp -= activmagic[a].atk
								if monsters[vismonsters[c].num].hp <= 0 {
									monsters[vismonsters[c].num].inactiv = true
									makededmonster(monsters[vismonsters[c].num].cnt)
								}
							}

						}
					}

				}

				// add flames
				if frames%6 == 0 {
					zfx := xfx{}
					v2 := findrandpoint(player.cnt)
					zfx.name = "flame"
					zfx.img = flameimg
					zfx.rec = rl.NewRectangle(v2.X, v2.Y, tilesize, tilesize)
					zfx.timer = rInt32(fpsint, fpsint*3)
					fx = append(fx, zfx)
				}

			}

		}

	}

}
func drawclearedlevel() { //MARK: drawclearedlevel

	txtlen := rl.MeasureText("level cleared", 100)
	rl.DrawText("level cleared", scrw/2-txtlen/2-3, scrh/2-80+3, 160, rl.White)
	rl.DrawText("level cleared", scrw/2-txtlen/2-2, scrh/2-80+2, 160, rl.Black)
	rl.DrawText("level cleared", scrw/2-txtlen/2, scrh/2-80, 160, randomcolor())

	if clearedlevellootnum > 0 {

		makenewobj(2, mousroomnum, player.cnt)
		newmsg("OH LOOK! lovely sparkly treasure to collect")

		if soundfxon && !mute {
			rl.PlaySoundMulti(gemaud)
		}

		clearedlevellootnum--

	}

}
func drawmsgs() { //MARK: drawmsgs

	x := int32(tilesize)
	y := int32(tilesize)

	count := len(msgs) - 1

	for {

		rl.DrawText(msgs[count], x, y, txtdef, rl.Black)
		y += 24

		count--
		if count <= 0 || y > scrh {
			break
		}
	}

}
func drawscore() { //MARK: drawscore

	rl.DrawRectangle(0, 0, scrw, scrh, rl.Black)

	scoreontimer++

	rl.DrawText("monsters killed", 100, 100, 40, brightorange())
	rl.DrawText(fmt.Sprint(monsterkills), 500, 100, 40, brightorange())

	rl.DrawText("bosses killed", 100, 150, 40, brightorange())
	rl.DrawText(fmt.Sprint(bosskills), 500, 150, 40, brightorange())

	rl.DrawText("dungeon level", 100, 200, 40, brightorange())
	rl.DrawText(fmt.Sprint(maxlevelreached), 500, 200, 40, brightorange())

	rl.DrawText("run time", 100, 250, 40, brightorange())

	runmintxt := fmt.Sprint(runmin)

	if runmin == 0 {
		runmintxt = "00"
	} else if runmin > 0 && runmin < 10 {
		runmintxt = "0" + fmt.Sprint(runmin)
	}

	runsectxt := fmt.Sprint(runsecs)

	if runsecs == 0 {
		runsectxt = "00"
	} else if runsecs > 0 && runsecs < 10 {
		runsectxt = "0" + fmt.Sprint(runsecs)
	}

	runtimetxt := runmintxt + ":" + runsectxt

	rl.DrawText(runtimetxt, 500, 250, 40, brightorange())

	rl.DrawText("score", 100, 350, 60, brightorange())
	rl.DrawText(fmt.Sprint(score), 500, 350, 60, brightorange())

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) && scoreontimer > 60 {
		scoreontimer = 0
		restart()
	}

	txtlen := rl.MeasureText("click left mouse button", 40)
	rl.DrawText("click left mouse button", scrw/2-(txtlen/2), scrh/2+150, 40, rl.Black)

	if rolldice() == 6 {
		rl.DrawText("click left mouse button", scrw/2-(txtlen/2), scrh/2+150, 40, randomcolor())
	}

	//close game
	closerec := rl.NewRectangle(scrwf32-(tilesize/2+tilesize/8), tilesize/8, tilesize/2, tilesize/2)
	rl.DrawRectangleRec(closerec, rl.Black)
	rl.DrawLine(closerec.ToInt32().X, closerec.ToInt32().Y, closerec.ToInt32().X+int32(tilesize/2), closerec.ToInt32().Y+int32(tilesize/2), brightorange())
	rl.DrawLine(closerec.ToInt32().X+int32(tilesize/2), closerec.ToInt32().Y, closerec.ToInt32().X, closerec.ToInt32().Y+int32(tilesize/2), brightorange())
	if rl.CheckCollisionPointRec(mousev2, closerec) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			pause = true
			endgamewindow = true
		}
	}

	//scanlines

	for a := 0; a < len(scanlines); a++ {
		rl.DrawLineV(scanlines[a].v1, scanlines[a].v2, rl.Fade(rl.Black, 0.7))
		scanlines[a].v1.Y += 1
		scanlines[a].v2.Y = scanlines[a].v1.Y

		if scanlines[a].v1.Y == scrhf32 {
			scanlines[a].v1.Y = 0
			scanlines[a].v2.Y = 0
		}
	}

}

// MARK: WEAPONS WEAPONS WEAPONS WEAPONS WEAPONS WEAPONS WEAPONS WEAPONS WEAPONS
func useweapon() { //MARK: useweapon

	zobj := xobj{}
	zobj.color = player.weapon.color
	zobj.atk = player.weapon.atk

	switch player.weapon.name {

	case "wand":
		if player.weapon.numberof > 0 {
			switch player.weapon.ability {
			case 4:
				if makelightning() {
					player.weapon.numberof--
					inven[activweaponnum].numberof--
				}
			case 3:
				if frograintimer == 0 {
					makefrogs()
				} else {
					newmsg("this wand has a cool down... wait a few seconds")
				}
				player.weapon.numberof--
				inven[activweaponnum].numberof--
			case 2:
				makeorbit()
				player.weapon.numberof--
				inven[activweaponnum].numberof--
			case 1:
				makelaser()
				player.weapon.numberof--
				inven[activweaponnum].numberof--
			}

		} else if player.weapon.numberof == 0 {
			inven[activweaponnum] = xobj{}
			if autoswitchweapons {
				if !findnextweapon() {
					player.weapon = xobj{}
					newmsg("no weapons in backpack... find some")
				}
			} else {
				player.weapon = xobj{}

			}
			findinvennum()

		}

	case "sword", "dagger", "mace", "club", "scythe", "axe":

		if len(boss) > 0 {

			for a := 0; a < len(boss); a++ {
				if boss[a].hppause == 0 {
					if rl.CheckCollisionRecs(boss[a].rec, player.weapon.meleerangerec) {
						chooseaud := monsterhit1aud
						choose := rInt(0, 7)
						switch choose {
						case 1:
							chooseaud = monsterhit2aud
						case 2:
							chooseaud = monsterhit3aud
						case 3:
							chooseaud = monsterhit4aud
						case 4:
							chooseaud = monsterhit5aud
						case 5:
							chooseaud = monsterhit6aud
						case 6:
							chooseaud = monsterhit7aud
						}
						if soundfxon && !mute {
							rl.PlaySoundMulti(chooseaud)
						}

						boss[a].hppause = fps / 2
						boss[a].hp -= player.weapon.atk + player.str
						if boss[a].hp <= 0 {
							boss[a].hp = 0
							boss[a].inactiv = true
							makeblood()
							bosskills++
							score += 100
							for b := 0; b < len(enemybullets); b++ {
								if enemybullets[b].bossbullet {
									enemybullets = remobj(enemybullets, b)
								}
							}
						}

						if player.vampirelev > 0 {

							choose := rolldice()

							if player.vampirelev == 1 {
								if choose > 4 {
									if player.hp < player.hpmax {
										player.hp++
										newmsg("hp+ vampirism...")
									}
								}
							} else if player.vampirelev == 2 {
								if choose > 2 {
									if player.hp < player.hpmax {
										player.hp++
										newmsg("hp+ vampirism...")
									}
								}

							} else {
								if player.hp < player.hpmax {
									player.hp++
									newmsg("hp+ vampirism...")
								}
							}

						}
					}
				}

			}

		}

		for a := 0; a < len(vismonsters); a++ {
			if vismonsters[a].hppause == 0 {
				if rl.CheckCollisionRecs(vismonsters[a].rec, player.weapon.meleerangerec) {

					chooseaud := monsterhit1aud
					choose := rInt(0, 7)
					switch choose {
					case 1:
						chooseaud = monsterhit2aud
					case 2:
						chooseaud = monsterhit3aud
					case 3:
						chooseaud = monsterhit4aud
					case 4:
						chooseaud = monsterhit5aud
					case 5:
						chooseaud = monsterhit6aud
					case 6:
						chooseaud = monsterhit7aud
					}
					if soundfxon && !mute {
						rl.PlaySoundMulti(chooseaud)
					}

					monsters[vismonsters[a].num].hppause = fps / 2
					vismonsters[a].hppause = fps / 2
					monsters[vismonsters[a].num].hp -= player.weapon.atk + player.str
					if player.vampirelev > 0 {

						choose := rolldice()

						if player.vampirelev == 1 {
							if choose > 4 {
								if player.hp < player.hpmax {
									player.hp++
									newmsg("hp+ vampirism...")
								}
							}
						} else if player.vampirelev == 2 {
							if choose > 2 {
								if player.hp < player.hpmax {
									player.hp++
									newmsg("hp+ vampirism...")
								}
							}

						} else {
							if player.hp < player.hpmax {
								player.hp++
								newmsg("hp+ vampirism...")
							}
						}

					}
					if monsters[vismonsters[a].num].hp <= 0 {
						monsters[vismonsters[a].num].inactiv = true
						makededmonster(monsters[vismonsters[a].num].cnt)
					}
				}
			}
		}

	case "ninja star", "throwing axe", "spear":
		if player.weapon.numberof > 0 {
			if mouseclicknum == 0 {
				zobj.dirx, zobj.diry = weapondir(player.weapon.vel1)
				zobj.img = player.weapon.img

				zobj.name = "rotates"
				if player.weapon.name == "spear" {
					zobj.name = ""
					if player.lr {
						zobj.ro = angle2points(player.cnt, weaponv2) - 45
					} else {
						zobj.ro = angle2points(player.cnt, weaponv2) + 45
					}
				}
				if selpoint.X > player.cnt.X { //player right
					zobj.rec = player.rec
					zobj.rec.Width = (tilesize / 4) * 3
					zobj.rec.Height = (tilesize / 4) * 3
					zobj.rec.X += tilesize + tilesize/5
					zobj.rec.Y += (zobj.rec.Height / 4) * 3
				} else {
					zobj.rec = player.rec
					zobj.rec.Width = (tilesize / 4) * 3
					zobj.rec.Height = (tilesize / 4) * 3
					zobj.rec.X -= tilesize / 4
					zobj.rec.Y += (zobj.rec.Height / 4) * 3
				}
				player.weapon.numberof--
				inven[activweaponnum].numberof--
				activweapons = append(activweapons, zobj)
			}
		} else if player.weapon.numberof == 0 {
			inven[activweaponnum] = xobj{}
			if autoswitchweapons {
				if !findnextweapon() {
					player.weapon = xobj{}
					newmsg("no weapons in backpack... find some")
				}
			} else {
				player.weapon = xobj{}
			}

			findinvennum()

		}
	case "crossbow", "bow":

		if player.ammo.numberof > 0 {
			if mouseclicknum == 0 {

				zobj.atk = player.ammo.atk + player.weapon.atk

				zobj.dirx, zobj.diry = weapondir(player.weapon.vel1)
				zobj.img = arrowimg
				zobj.ro = angle2points(player.cnt, weaponv2) - 45

				if selpoint.X > player.cnt.X { //player right
					zobj.rec = player.rec
					zobj.rec.Width = tilesize / 2
					zobj.rec.Height = tilesize / 2
					zobj.rec.X += tilesize + tilesize/5
					zobj.rec.Y += (zobj.rec.Height / 4) * 3
				} else {
					zobj.rec = player.rec
					zobj.rec.Width = tilesize / 2
					zobj.rec.Height = tilesize / 2
					zobj.rec.X -= tilesize / 4
					zobj.rec.Y += (zobj.rec.Height / 4) * 3
				}
				player.ammo.numberof--
				inven[activammonum].numberof--
				activweapons = append(activweapons, zobj)

			}
		} else {
			if activammonum != blankint {
				inven[activammonum] = xobj{}
				player.ammo = xobj{}
			}

			if autoswitchammo {
				if !findnextammo() {
					newmsg("NO AMMO >> check inventory or find some")
					if autoswitchweapons {
						if !findnextweapon() {
							newmsg("no other weapons in backpack... find some")
						} else {

						}
					}
				}
			} else {
				newmsg("NO AMMO >> check inventory or find some")
				inven[activammonum] = xobj{}
				activammonum = blankint
				player.ammo = xobj{}
				findinvennum()

			}
		}
	}

}
func weapondir(vel float32) (dirx, diry float32) { //MARK: weapondir

	x, y := float32(0), float32(0)

	xdiff := absdiff32(player.cnt.X, weaponv2.X)
	ydiff := absdiff32(player.cnt.Y, weaponv2.Y)

	if xdiff > ydiff {
		x = vel
		if weaponv2.X < player.cnt.X {
			x = -x
		}
		ychange := xdiff / vel
		y = ydiff / ychange
		if weaponv2.Y < player.cnt.Y {
			y = -y
		}
	} else {
		y = vel
		if weaponv2.Y < player.cnt.Y {
			y = -y
		}
		xchange := ydiff / vel
		x = xdiff / xchange
		if weaponv2.X < player.cnt.X {
			x = -x
		}

	}

	return x, y

}
func checkweaponcollisions() { //MARK: checkweaponcollisions

	for a := 0; a < len(activweapons); a++ {

		if len(boss) > 0 {
			for b := 0; b < len(boss); b++ {
				if !boss[b].inactiv {
					if activweapons[a].noimg {

						if boss[b].hppause == 0 {
							if rl.CheckCollisionPointRec(activweapons[a].v1, boss[b].rec) || rl.CheckCollisionPointRec(activweapons[a].v2, boss[b].rec) {

								chooseaud := monsterhit1aud
								choose := rInt(0, 7)
								switch choose {
								case 1:
									chooseaud = monsterhit2aud
								case 2:
									chooseaud = monsterhit3aud
								case 3:
									chooseaud = monsterhit4aud
								case 4:
									chooseaud = monsterhit5aud
								case 5:
									chooseaud = monsterhit6aud
								case 6:
									chooseaud = monsterhit7aud
								}
								if soundfxon && !mute {
									rl.PlaySoundMulti(chooseaud)
								}

								boss[b].hppause = fps / 2
								boss[b].hp -= activweapons[a].atk + player.intel

								if boss[b].hp <= 0 {
									boss[b].hp = 0
									boss[a].inactiv = true
									makeblood()
									bosskills++
									score += 100
									for c := 0; c < len(enemybullets); c++ {
										if enemybullets[c].bossbullet {
											enemybullets = remobj(enemybullets, c)
										}
									}
								}

								if player.vampirelev > 0 {
									choose := rolldice()
									if player.vampirelev == 1 {
										if choose > 4 {
											if player.hp < player.hpmax {
												player.hp++
												newmsg("hp+ vampirism...")
											}
										}
									} else if player.vampirelev == 2 {
										if choose > 2 {
											if player.hp < player.hpmax {
												player.hp++
												newmsg("hp+ vampirism...")
											}
										}

									} else {
										if player.hp < player.hpmax {
											player.hp++
											newmsg("hp+ vampirism...")
										}
									}
								}
							}

						}

					} else {

						if rl.CheckCollisionRecs(activweapons[a].rec, boss[b].rec) {

							if activweapons[a].usetype != 3 {
								activweapons[a].color = randomcolor()
							}

							chooseaud := monsterhit1aud
							choose := rInt(0, 7)
							switch choose {
							case 1:
								chooseaud = monsterhit2aud
							case 2:
								chooseaud = monsterhit3aud
							case 3:
								chooseaud = monsterhit4aud
							case 4:
								chooseaud = monsterhit5aud
							case 5:
								chooseaud = monsterhit6aud
							case 6:
								chooseaud = monsterhit7aud
							}
							if soundfxon && !mute {
								rl.PlaySoundMulti(chooseaud)
							}

							boss[b].hppause = fps / 2
							boss[b].hp -= activweapons[a].atk + player.dex
							if boss[b].hp <= 0 {
								boss[b].hp = 0
								boss[b].inactiv = true
								makeblood()
							}

							if player.vampirelev > 0 {
								choose := rolldice()
								if player.vampirelev == 1 {
									if choose > 4 {
										if player.hp < player.hpmax {
											player.hp++
											newmsg("hp+ vampirism...")
										}
									}
								} else if player.vampirelev == 2 {
									if choose > 2 {
										if player.hp < player.hpmax {
											player.hp++
											newmsg("hp+ vampirism...")
										}
									}

								} else {
									if player.hp < player.hpmax {
										player.hp++
										newmsg("hp+ vampirism...")
									}
								}
							}

						}
					}
				}
			}
		}

		for b := 0; b < len(vismonsters); b++ {

			if activweapons[a].noimg {

				if vismonsters[b].hppause == 0 {
					if rl.CheckCollisionPointRec(activweapons[a].v1, vismonsters[b].rec) || rl.CheckCollisionPointRec(activweapons[a].v2, vismonsters[b].rec) {

						chooseaud := monsterhit1aud
						choose := rInt(0, 7)
						switch choose {
						case 1:
							chooseaud = monsterhit2aud
						case 2:
							chooseaud = monsterhit3aud
						case 3:
							chooseaud = monsterhit4aud
						case 4:
							chooseaud = monsterhit5aud
						case 5:
							chooseaud = monsterhit6aud
						case 6:
							chooseaud = monsterhit7aud
						}
						if soundfxon && !mute {
							rl.PlaySoundMulti(chooseaud)
						}
						monsters[vismonsters[b].num].hppause = fps / 2
						vismonsters[b].hppause = fps / 2
						monsters[vismonsters[b].num].hp -= activweapons[a].atk + player.intel
						if player.vampirelev > 0 {
							choose := rolldice()
							if player.vampirelev == 1 {
								if choose > 4 {
									if player.hp < player.hpmax {
										player.hp++
										newmsg("hp+ vampirism...")
									}
								}
							} else if player.vampirelev == 2 {
								if choose > 2 {
									if player.hp < player.hpmax {
										player.hp++
										newmsg("hp+ vampirism...")
									}
								}

							} else {
								if player.hp < player.hpmax {
									player.hp++
									newmsg("hp+ vampirism...")
								}
							}
						}
						if monsters[vismonsters[b].num].hp <= 0 {
							monsters[vismonsters[b].num].inactiv = true
							makededmonster(monsters[vismonsters[b].num].cnt)
						}
					}

				}

			} else {
				if rl.CheckCollisionRecs(activweapons[a].rec, vismonsters[b].rec) {

					if activweapons[a].usetype != 3 {
						activweapons[a].color = randomcolor()
					}
					if vismonsters[b].hppause == 0 {
						chooseaud := monsterhit1aud
						choose := rInt(0, 7)
						switch choose {
						case 1:
							chooseaud = monsterhit2aud
						case 2:
							chooseaud = monsterhit3aud
						case 3:
							chooseaud = monsterhit4aud
						case 4:
							chooseaud = monsterhit5aud
						case 5:
							chooseaud = monsterhit6aud
						case 6:
							chooseaud = monsterhit7aud
						}
						if soundfxon && !mute {
							rl.PlaySoundMulti(chooseaud)
						}
						monsters[vismonsters[b].num].hppause = fps / 2
						vismonsters[b].hppause = fps / 2
						monsters[vismonsters[b].num].hp -= activweapons[a].atk + player.dex
						if player.vampirelev > 0 {
							choose := rolldice()
							if player.vampirelev == 1 {
								if choose > 4 {
									if player.hp < player.hpmax {
										player.hp++
										newmsg("hp+ vampirism...")
									}
								}
							} else if player.vampirelev == 2 {
								if choose > 2 {
									if player.hp < player.hpmax {
										player.hp++
										newmsg("hp+ vampirism...")
									}
								}

							} else {
								if player.hp < player.hpmax {
									player.hp++
									newmsg("hp+ vampirism...")
								}
							}
						}
						if monsters[vismonsters[b].num].hp <= 0 {
							monsters[vismonsters[b].num].inactiv = true
							makededmonster(monsters[vismonsters[b].num].cnt)
						}
					}
				}
			}
		}

	}

}

// MARK: OBJS OBJS OBJS OBJS OBJS OBJS OBJS OBJS OBJS OBJS OBJS OBJS OBJS OBJS
func objcollisionactions(roomnum, objnum int) { //MARK: objcollisionactions

	switch level[roomnum].objs[objnum].name {

	case "spike trap":
		if player.hpppause == 0 && !spikeimmunity {
			if soundfxon && !mute {
				rl.PlaySoundMulti(spikesaud)
			}
			player.hp -= 3 + currentlevelnum
			if player.hp < 0 {
				player.hp = 0
			}
			player.hpppause = fps
			makeblood()
			newmsg("you stood on a spike trap and it hurt quite a bit...")
		} else if spikeimmunity {
			newmsg("you are immune to spike trap damage this level...")
		}
	case "spring":
		springv2 = level[roomnum].objs[objnum].v1
		player.nomove = true
		newmsg("you have bounced to another part of the dungeon...")
		if soundfxon && !mute {
			rl.PlaySoundMulti(springaud)
		}

	case "cactus":
		if player.hpppause == 0 && !cactusimmunity {
			player.hp -= 1 + currentlevelnum
			if soundfxon && !mute {
				rl.PlaySoundMulti(cactusaud)
			}
			if player.hp < 0 {
				player.hp = 0
			}
			player.hpppause = fps
			newmsg("OUCH! cactus thorns damage...")
		} else if cactusimmunity {
			newmsg("you are immune to cactus damage this level...")
		}

	case "water":
		if player.burning {
			player.burning = false
			player.burntimer = 0
			newmsg("YAY >> no longer burning")
		} else {
			if !level[roomnum].objs[objnum].msgadded {
				newmsg("be careful... you might catch a cold")
				level[roomnum].objs[objnum].msgadded = true
			}
		}
		if player.damppause == 0 && !player.immune {
			player.damptimer = fps * 4
			player.damppause = fps * 2
			player.dampcount++
			if player.dampcount > 2 {
				player.dampcount = 2
				player.sick = true
				player.sicktimer = fps * 4
				newmsg("you and are now SICK")
			}
		}

	}

}
func clickobj(roomnum, objnum int) { //MARK: clickobj

	zobj := level[roomnum].objs[objnum]

	switch zobj.name {
	case "ornamental plant":
		if !level[roomnum].objs[objnum].inactiv {
			num := rInt(3, 8)
			player.teleports += num
			if player.teleports > 99 {
				player.teleports = 99
			}
			newmsg("you are suddenly covered in magical pixie dust... teleports +" + fmt.Sprint(num))
			if soundfxon && !mute {
				rl.PlaySoundMulti(teleportpickupaud)
			}
			level[roomnum].objs[objnum].inactiv = true
		} else {
			newmsg("you have already made use of this object - it is now inactive")
		}
	case "magic watering can":
		if !level[roomnum].objs[objnum].inactiv {
			if player.coins < refillcost {
				newmsg("not enough coins... make no bones about it (that is a badly disguised hint)")
			} else {
				player.coins -= refillcost
				player.hpmax += rInt(1, 5)
				player.hp = player.hpmax
				player.str += rInt(1, 5)
				player.luck += rInt(1, 5)
				player.intel += rInt(1, 5)
				player.dex += rInt(1, 5)
				newmsg("you took a resreshing shower using a rusty watering can... all stats increased")
				level[roomnum].objs[objnum].inactiv = true
			}
		} else {
			newmsg("you have already made use of this object - it is now inactive")
		}
	case "cauldron":
		makepotions()
	case "weed":
		if invencurrentnum < len(inven) {
			inven[invencurrentnum] = level[roomnum].objs[objnum]
			invencurrentnum++
			level[roomnum].objs[objnum].nodraw = true
			newmsg("i wonder what i could do with this...")
			if soundfxon && !mute {
				rl.PlaySoundMulti(weedaud)
			}
		}
	case "staircase":
		if soundfxon && !mute {
			rl.PlaySoundMulti(stairsaud)
		}
		level[roomnum].objs[objnum].inactiv = true
		changelevelon = true

	case "bookcase":
		if !level[roomnum].objs[objnum].inactiv {
			player.intel++
			level[roomnum].objs[objnum].inactiv = true
			newmsg("read a book >> intelligence +1")
			if soundfxon && !mute {
				rl.PlaySoundMulti(bookcaseaud)
			}
		} else {
			newmsg("you have already made use of this object - it is now inactive")
		}

	case "chest":

		if zobj.inactiv {
			num := rInt(1, 3)

			for {
				makenewobj(1, roomnum, zobj.cnt)
				num--
				if num == 0 {
					break
				}
			}
			if soundfxon && !mute {
				rl.PlaySoundMulti(chestaud)
			}

			newmsg("opened a chest")
		} else {
			newimg := chestimg
			newimg.X += 24
			level[roomnum].objs[objnum].img = newimg
			level[roomnum].objs[objnum].inactiv = true
		}

	}

}

func objactions(roomnum, objnum int) { //MARK: objactions

	if !level[roomnum].objs[objnum].fixed {
		level[roomnum].objs[objnum].ro += 1
	}

	switch level[roomnum].objs[objnum].name {
	case "spring":
		rl.DrawCircleGradient(int32(level[roomnum].objs[objnum].cnt.X), int32(level[roomnum].objs[objnum].cnt.Y), rFloat32(30, 60), rl.Fade(rl.Magenta, rFloat32(0.3, 0.7)), rl.Blank)
	case "spike trap":

		if level[roomnum].objs[objnum].onoff {
			level[roomnum].objs[objnum].img.Y++
			if level[roomnum].objs[objnum].img.Y >= spikeimg.Y {
				level[roomnum].objs[objnum].onoff = false
			}

		} else {
			level[roomnum].objs[objnum].img.Y--
			if level[roomnum].objs[objnum].img.Y <= spikeimg.Y-spikeimg.Height {
				level[roomnum].objs[objnum].onoff = true
			}
		}

	case "chest":
		if !level[roomnum].objs[objnum].inactiv {
			rl.DrawCircleGradient(int32(level[roomnum].objs[objnum].cnt.X), int32(level[roomnum].objs[objnum].cnt.Y), rFloat32(30, 60), rl.Fade(brightyellow(), rFloat32(0.3, 0.7)), rl.Blank)
		} else if level[roomnum].objs[objnum].inactiv && !level[roomnum].objs[objnum].genericswitch {
			if frames%20 == 0 {
				newimg := chestimg
				newimg.X += 48
				level[roomnum].objs[objnum].img = newimg
				level[roomnum].objs[objnum].genericswitch = true
				clickobj(roomnum, objnum)
			}

		}
	case "weed", "ornamental plant", "cactus", "gem", "jewelry", "legendary armor":
		rl.DrawCircleGradient(int32(level[roomnum].objs[objnum].cnt.X), int32(level[roomnum].objs[objnum].cnt.Y), rFloat32(30, 60), rl.Fade(level[roomnum].objs[objnum].color, rFloat32(0.3, 0.7)), rl.Blank)
	case "food":
		rl.DrawCircleGradient(int32(level[roomnum].objs[objnum].cnt.X), int32(level[roomnum].objs[objnum].cnt.Y), rFloat32(30, 60), rl.Fade(brightred(), rFloat32(0.3, 0.7)), rl.Blank)
	case "torch":
		rl.DrawCircleGradient(int32(level[roomnum].objs[objnum].cnt.X), int32(level[roomnum].objs[objnum].cnt.Y), rFloat32(40, 120), rl.Fade(brightyellow(), rFloat32(0.3, 0.7)), rl.Blank)
	case "camp fire":
		level[roomnum].objs[objnum].color = randomorange()
		rl.DrawCircleGradient(int32(level[roomnum].objs[objnum].cnt.X), int32(level[roomnum].objs[objnum].cnt.Y), rFloat32(30, 80), rl.Fade(randomorange(), rFloat32(0.4, 0.9)), rl.Blank)

		if rl.CheckCollisionRecs(player.boundrec, level[roomnum].objs[objnum].rec) {
			if player.hp != player.hpmax {
				if frames%int(fps) == 0 {
					player.hp++
					newmsg("It is so cosy here >> HP +1")
					if soundfxon && !mute {
						rl.PlaySoundMulti(campaud)
					}

				}
			}
		}

	}

}

func objboundreccollisionactions(roomnum, objnum int) { //MARK: objboundreccollisionactions

	switch level[roomnum].objs[objnum].name {
	case "spring":

		if level[roomnum].objs[objnum].onoff {

			if level[roomnum].objs[objnum].img.X == springimg.X+springimg.Width*2 {
				if frames%10 == 0 {
					level[roomnum].objs[objnum].img.X -= springimg.Width
				}
			} else if level[roomnum].objs[objnum].img.X == springimg.X+springimg.Width {
				if frames%10 == 0 {
					level[roomnum].objs[objnum].img.X -= springimg.Width
				}
			} else if level[roomnum].objs[objnum].img.X == springimg.X {

				level[roomnum].objs[objnum].onoff = false

			}

		} else {

			if level[roomnum].objs[objnum].img.X == springimg.X {
				if frames%10 == 0 {
					level[roomnum].objs[objnum].img.X += springimg.Width
				}
			} else if level[roomnum].objs[objnum].img.X == springimg.X+springimg.Width {
				if frames%10 == 0 {
					level[roomnum].objs[objnum].img.X += springimg.Width
				}
			} else if level[roomnum].objs[objnum].img.X == springimg.X+springimg.Width*2 {
				level[roomnum].objs[objnum].onoff = true
			}

		}

	case "ornamental plant":
		if !level[roomnum].objs[objnum].inactiv && !level[roomnum].objs[objnum].msgadded {
			newmsg("there is something mystical about this plant perhaps...")
			level[roomnum].objs[objnum].msgadded = true
		}
	case "old bones":
		if !level[roomnum].objs[objnum].inactiv {
			if !level[roomnum].objs[objnum].msgadded {
				newmsg("i wonder if he was a rich man...")
				level[roomnum].objs[objnum].msgadded = true
			}
			if rl.IsMouseButtonPressed(rl.MouseRightButton) {
				if player.weapon.name != "" {
					if rolldice() > 3 {
						makenewobj(3, roomnum, player.cnt)
						newmsg("it would appear that he was... you have found a coin")
						level[roomnum].objs[objnum].nodraw = true

					}
				}
				level[roomnum].objs[objnum].inactiv = true
			}
		}
	case "cactus":
		if !level[roomnum].objs[objnum].msgadded {
			newmsg("be careful... i have a prickly feeling...")
		}
		level[roomnum].objs[objnum].inbound = true
	case "grave":
		if !level[roomnum].objs[objnum].inactiv {
			if !level[roomnum].objs[objnum].msgadded {
				newmsg("the ground here is soft... i wonder what is beneath")
				level[roomnum].objs[objnum].msgadded = true
			}

			if rl.IsMouseButtonPressed(rl.MouseRightButton) && player.object.name == "spade" {
				if !level[roomnum].objs[objnum].inactiv {
					choose := rInt(1, 101)

					if choose <= player.luck {
						makenewobj(2, roomnum, level[roomnum].objs[objnum].cnt)
						newmsg("OH LOOK! lovely sparkly treasure to collect")
						level[roomnum].objs[objnum].inactiv = true
						if soundfxon && !mute {
							rl.PlaySoundMulti(gemaud)
						}
					}
				}
			}
		}
	}

}
func collectobj(roomnum, objnum int) { //MARK: collectobj

	if level[roomnum].objs[objnum].name == "food" {
		if player.hp < player.hpmax {
			player.hp++
			newmsg("YUM! food >> HP +1")
			level[roomnum].objs[objnum].nodraw = true
			if soundfxon && !mute {
				rl.PlaySoundMulti(eataud)
			}
		}

	} else if level[roomnum].objs[objnum].name == "coin" {

		player.coins++
		newmsg("you are now officially one coin richer...")

		if soundfxon && !mute {
			rl.PlaySoundMulti(collectaud)

		}

		level[roomnum].objs[objnum].nodraw = true
	} else if level[roomnum].objs[objnum].kind == "scroll" {
		if beltinvencurrentnum < len(beltinven) {
			beltinven[beltinvencurrentnum] = level[roomnum].objs[objnum]
			level[roomnum].objs[objnum].nodraw = true
			findbeltinvennum()
			if soundfxon && !mute {
				rl.PlaySoundMulti(collectaud)
			}
		} else {
			if invencurrentnum < len(inven) {
				inven[invencurrentnum] = level[roomnum].objs[objnum]
				level[roomnum].objs[objnum].nodraw = true
				findinvennum()
				if soundfxon && !mute {
					rl.PlaySoundMulti(collectaud)
				}
			} else {
				newmsg("belt inventory is full, right click to destroy or move items")
			}
		}

	} else {

		if invencurrentnum < len(inven) {
			inven[invencurrentnum] = level[roomnum].objs[objnum]
			level[roomnum].objs[objnum].nodraw = true

			if level[roomnum].objs[objnum].questitem {
				questitemon = false
				questitemv2 = blankv2
				newmsg("congratulations on acquiring a sparkly new legendary item, remember these are part of a set")
			}
			findinvennum()
			if soundfxon && !mute {
				rl.PlaySoundMulti(collectaud)
			}
		} else {
			if level[roomnum].objs[objnum].kind == "scroll" || level[roomnum].objs[objnum].kind == "key" || level[roomnum].objs[objnum].kind == "map" || level[roomnum].objs[objnum].kind == "gem" {
				if beltinvencurrentnum < len(beltinven) {
					beltinven[beltinvencurrentnum] = level[roomnum].objs[objnum]
					level[roomnum].objs[objnum].nodraw = true
					findbeltinvennum()
					if soundfxon && !mute {
						rl.PlaySoundMulti(collectaud)
					}
				} else {
					newmsg("backpack inventory is full, right click to destroy or move items")
				}

			} else {
				newmsg("backpack inventory is full, right click to destroy or move items")
			}
		}
	}

}

//MARK: FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND

func findrandpointinroom(cnt rl.Vector2) rl.Vector2 { //MARK: findrandpointinroom

	v2 := cnt

	found := false
	breakcount := 100
	for {
		v2.X += rFloat32(-tilesize*8, tilesize*8)
		v2.Y += rFloat32(-tilesize*8, tilesize*8)

		for a := 0; a < len(visroom); a++ {

			for b := 0; b < len(visroom[a].roomrec); b++ {
				if rl.CheckCollisionPointRec(v2, visroom[a].roomrec[b].rec) {
					found = true
					break
				}

			}
			if found {
				break
			}

		}

		breakcount--
		if breakcount == 0 || found {
			break
		}

	}

	return v2

}

func findinvennum() { //MARK: findinvennum

	for a := 0; a < len(inven); a++ {

		if inven[a].name == "" {
			invencurrentnum = a
			break
		} else {
			invencurrentnum = len(inven)
		}

	}

}
func findbeltinvennum() { //MARK: findbeltinvennum

	for a := 0; a < len(beltinven); a++ {

		if beltinven[a].name == "" {
			beltinvencurrentnum = a
			break
		} else {
			beltinvencurrentnum = len(beltinven)
		}

	}

}
func findcntr() (rl.Vector2, int) { //MARK: findcntr

	roomnum := 0
	v2 := rl.NewVector2(0, 0)

	for {

		x := rFloat32(levboundrec.X, levboundrec.X+levboundrec.Width)
		y := rFloat32(levboundrec.Y, levboundrec.Y+levboundrec.Height)
		v2 = rl.NewVector2(x, y)

		inlevel := false

		for a := 0; a < len(level); a++ {

			for b := 0; b < len(level[a].roomrec); b++ {

				if rl.CheckCollisionPointRec(v2, level[a].roomrec[b].innerrec) {

					roomnum = a
					inlevel = true

				}

			}

		}

		if inlevel {

			for a := 0; a < len(level[roomnum].objs); a++ {

				if rl.CheckCollisionPointRec(v2, level[roomnum].objs[a].boundrec) {
					inlevel = false
				}

			}

		}

		if inlevel {

			break
		}
	}

	return v2, roomnum
}
func findrandpoint(cnt rl.Vector2) rl.Vector2 { //MARK: findrandpoint

	v2 := cnt

	v2.X += rFloat32(-tilesize*4, tilesize*4)
	v2.Y += rFloat32(-tilesize*4, tilesize*4)

	return v2

}
func findvismonsternum(num int) int { //MARK: findvismonsternum

	monsternum := 0

	for a := 0; a < len(vismonsters); a++ {
		if vismonsters[a].num == num {
			monsternum = a
		}
	}
	return monsternum
}
func findnextweapon() bool { //MARK: findnextweapon

	found := false
	switch player.weapon.name {
	case "wand":
		for a := 0; a < len(inven); a++ {
			if inven[a].name == "wand" {
				inven[a].invenselect = false
				activweaponnum = a
				player.weapon = inven[a]
				inven[a].invenselect = true
				found = true
				break
			}
		}
	}

	if !found {
		for a := 0; a < len(inven); a++ {
			switch inven[a].name {
			case "ninja star", "throwing axe", "spear":
				inven[activweaponnum].invenselect = false
				activweaponnum = a
				player.weapon = inven[a]
				inven[a].invenselect = true
				found = true
				break
			}
		}
	}

	if !found {

		found2 := false
		num2 := 0
		found3 := false
		num3 := 0

		for a := 0; a < len(inven); a++ {
			switch inven[a].name {
			case "ammo":
				found2 = true
				num2 = a
				break
			}
		}
		if found2 {
			for a := 0; a < len(inven); a++ {
				switch inven[a].name {
				case "bow", "crossbow":
					found3 = true
					num3 = a
					break
				}
			}
		}

		if found2 && found3 {
			player.weapon = inven[num3]
			inven[num3].invenselect = true
			inven[activweaponnum].invenselect = false
			activweaponnum = num3
			player.ammo = inven[num2]
			inven[num2].invenselect = true
			found = true
		}

	}

	if !found {

		for a := 0; a < len(inven); a++ {
			switch inven[a].name {
			case "sword", "dagger", "mace", "club", "scythe", "axe":
				inven[activweaponnum].invenselect = false
				activweaponnum = a
				player.weapon = inven[a]
				inven[a].invenselect = true
				found = true
				break
			}
		}

	}

	return found

}
func findnextammo() bool { //MARK: findnextammo

	found := false

	for a := 0; a < len(inven); a++ {
		if inven[a].kind == "ammo" {
			player.ammo = inven[a]
			activammonum = a
			inven[a].invenselect = true
			found = true
			break
		}
	}
	return found
}
func finrandpointborderrec() rl.Vector2 { //MARK: finrandpointborderrec

	point := rl.NewVector2(rFloat32(0, scrwf32), rFloat32(0, scrhf32))

	pointworld := rl.GetScreenToWorld2D(point, camera)

	return pointworld

}

// MARK: CLEAR CLEAR CLEAR CLEAR CLEAR CLEAR CLEAR CLEAR CLEAR CLEAR CLEAR
func clearenemybullets() { //MARK: clearenemybullets

	for a := 0; a < len(enemybullets); a++ {
		if enemybullets[a].inactiv {
			enemybullets = remobj(enemybullets, a)
		}
	}

}
func clearactivweapons() { //MARK: clearactivweapons

	for a := 0; a < len(activweapons); a++ {
		if activweapons[a].inactiv {
			activweapons = remobj(activweapons, a)
		}
	}

}

func cleanlevel() { //MARK: cleanlevel

}

// MARK: MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE
func makeboss() { //MARK: makeboss

	zboss := xboss{}
	zboss.rec.Width = tilesize * 4
	zboss.rec.Height = tilesize * 4
	zboss.cnt = findrandpointinroom(player.cnt)
	zboss.rec.X = zboss.cnt.X - zboss.rec.Width/2
	zboss.rec.Y = zboss.cnt.Y - zboss.rec.Height/2
	zboss.vel = 4
	zboss.follow = true
	zboss.attacktype = rInt(0, 4)

	zboss.hp = currentlevelnum * rInt(5, 9)

	zboss.atk = rInt(2, 5)
	zboss.atkspeed = fpsint / 2
	zboss.num = 0

	choose := rInt(1, 5)
	switch choose {
	case 1:
		zboss.bulletimg = bullet1img
	case 2:
		zboss.bulletimg = bullet2img
	case 3:
		zboss.bulletimg = bullet3img
	case 4:
		zboss.bulletimg = bullet4img
	}

	choose = rInt(0, 9)

	switch choose {
	case 8:
		zboss.name = "mr radish"
		zboss.img = radishboss[0]
	case 7:
		zboss.name = "mr spike"
		zboss.img = spikeboss[0]
	case 6:
		zboss.name = "mr skull"
		zboss.img = skullboss[0]
	case 5:
		zboss.name = "mr ghost"
		zboss.img = ghostboss[0]
	case 4:
		zboss.name = "mr reaper"
		zboss.img = reaperboss[0]
	case 3:
		zboss.name = "mr orc"
		zboss.img = orcboss[0]
	case 2:
		zboss.name = "mr slime"
		zboss.img = slimeboss[0]
	case 1:
		zboss.name = "mr dino"
		zboss.img = dinoboss[0]
	case 0:
		zboss.name = "mr mushroom"
		zboss.img = mushroomboss[0]
	}

	boss = append(boss, zboss)
	newmsg(zboss.name + " has arrived")
}

func makeshop() { //MARK: makeshop

	shopitems = nil
	shoprecs = nil

	num := 5

	length := tilesize * 5
	x := scrwf32 + tilesize
	y := scrhf32/2 - length/2

	shoprecx = x

	//make recs
	for a := 0; a < num; a++ {
		shoprecs = append(shoprecs, rl.NewRectangle(x, y, length, length))
		x += length + tilesize*2
	}

	shoprectotallen = length*float32(num) + ((tilesize * 2) * float32(num))

	//make items
	for a := 0; a < num; a++ {

		zitem := xshopitem{}

		choose := rInt(0, 16)

		switch choose {
		case 0:
			zitem.amount = rInt(5, 16)
			zitem.name = "hp max +" + fmt.Sprint(zitem.amount)
			zitem.name2 = "hpmax"
			zitem.img = hpmaximg
			zitem.color = rl.White
			zitem.cost = gemstotal
		case 1:
			zitem.amount = rInt(5, 16)
			zitem.name = "strength +" + fmt.Sprint(zitem.amount)
			zitem.name2 = "str"
			zitem.img = strimg
			zitem.color = rl.White
			zitem.cost = rInt(gemstotal/2, gemstotal+1)
		case 2:
			zitem.amount = rInt(5, 16)
			zitem.name = "intelligence +" + fmt.Sprint(zitem.amount)
			zitem.name2 = "int"
			zitem.img = intimg
			zitem.color = rl.White
			zitem.cost = rInt(gemstotal/2, gemstotal+1)
		case 3:
			zitem.amount = rInt(5, 16)
			zitem.name = "dexterity +" + fmt.Sprint(zitem.amount)
			zitem.name2 = "dex"
			zitem.img = deximg
			zitem.color = rl.White
			zitem.cost = rInt(gemstotal/2, gemstotal+1)
		case 4:
			zitem.amount = rInt(5, 16)
			zitem.name = "luck +" + fmt.Sprint(zitem.amount)
			zitem.name2 = "luk"
			zitem.img = lukimg
			zitem.color = rl.White
			zitem.cost = rInt(gemstotal/2, gemstotal+1)
		case 5:
			zitem.name = "fire resistance next level"
			zitem.name2 = "fireresist"
			zitem.img = fireresistimg
			zitem.color = randomorange()
			zitem.cost = rInt(gemstotal/3, gemstotal+1)
		case 6:
			zitem.name = "poison resistance next level"
			zitem.name2 = "poison"
			zitem.img = poisonresistimg
			zitem.color = randomgreen()
			zitem.cost = rInt(gemstotal/3, gemstotal+1)
		case 7:
			zitem.name = "teleports = 99"
			zitem.name2 = "tel99"
			zitem.img = teleportimg
			zitem.cost = gemstotal
		case 8:
			zitem.amount = rInt(5, 16)
			zitem.name = "coins +" + fmt.Sprint(zitem.amount)
			zitem.name2 = "coins"
			zitem.img = coinimg
			zitem.cost = rInt(gemstotal/3, gemstotal+1)
		case 9:
			zitem.name = "fire trail next level"
			zitem.name2 = "firetrail"
			zitem.img = flameimg
			zitem.color = randomorange()
			zitem.cost = rInt(gemstotal/4, gemstotal+1)
		case 10:
			zitem.name = "refill hp"
			zitem.name2 = "hp"
			zitem.img = hpimg
			zitem.color = brightred()
			zitem.cost = rInt(gemstotal/4, (gemstotal/2)+1)
		case 11:
			zitem.name = "random legendary armor at start of next level"
			zitem.name2 = "armor"
			zitem.img = vestimgs[rInt(0, len(vestimgs))]
			zitem.color = brightyellow()
			zitem.cost = rInt(gemstotal/2, gemstotal+1)
		case 12:
			zitem.name = "cactus immunity"
			zitem.name2 = "cactus"
			zitem.img = cactusimg
			zitem.color = randomgreen()
			zitem.cost = rInt(gemstotal/4, (gemstotal/2)+1)
		case 13:
			zitem.name = "spike trap immunity"
			zitem.name2 = "spike"
			zitem.img = spikeimg
			zitem.color = randombluelight()
			zitem.cost = rInt(gemstotal/4, (gemstotal/2)+1)
		case 14:
			zitem.name = "disease immunity"
			zitem.name2 = "disease"
			zitem.img = diseaseimmuneimg
			zitem.color = randomyellow()
			zitem.cost = rInt(gemstotal/4, (gemstotal/2)+1)
		case 15:
			zitem.amount = rInt(5, 16)
			zitem.name = "teleports +" + fmt.Sprint(zitem.amount)
			zitem.name2 = "tel"
			zitem.img = teleportimg
			zitem.color = rl.White
			zitem.cost = rInt(gemstotal/3, gemstotal+1)
		}

		if a == 3 || a == 4 {
			zitem.locked = true
		}

		shopitems = append(shopitems, zitem)
	}

	shopgemimg = gemimgs[rInt(0, len(gemimgs))]
	choosekey := rInt(1, 4)
	shopkeyimg = key1img
	switch choosekey {
	case 2:
		shopkeyimg = key2img
	case 3:
		shopkeyimg = key3img
	}

}

func makestartobjs() { //MARK: makestartobjs

	/*
		zobj := xobj{}
		zobj.cnt = player.cnt
		zobj.cnt.X += tilesize * 2

		choose := rInt(0, len(jewelryimgs))
		zobj.img = jewelryimgs[choose]
		zobj.color = brightyellow()
		zobj.collect = true
		zobj.name = "jewelry"
		zobj.amount = rInt(5, 21)
		zobj.usetype = 9
		zobj.kind = "jewel"
		switch zobj.usetype {
		case 1:
			zobj.name2 = "jewelry of dexterity +" + fmt.Sprint(zobj.amount)
		case 2:
			zobj.name2 = "jewelry of strength +" + fmt.Sprint(zobj.amount)
		case 3:
			zobj.name2 = "jewelry of luck +" + fmt.Sprint(zobj.amount)
		case 4:
			zobj.name2 = "jewelry of intelligence +" + fmt.Sprint(zobj.amount)
		case 5:
			zobj.amount = rInt(1, 11)
			zobj.name2 = "jewelry of fire protection +" + fmt.Sprint(zobj.amount*10) + "%"
		case 6:
			zobj.amount = rInt(1, 11)
			zobj.name2 = "jewelry of poison protection +" + fmt.Sprint(zobj.amount*10) + "%"
		case 7:
			zobj.name2 = "jewelry of disease immunity"
		case 8:
			zobj.name2 = "jewelry of max hp +" + fmt.Sprint(zobj.amount)
		case 9:
			zobj.name2 = "jewelry of teleport"
			zobj.timer = rInt32(fpsint*5, fpsint*10)
		}

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)

		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4

		level[0].objs = append(level[0].objs, zobj)
	*/

	//chests
	for a := 1; a < origlevellen; a++ {

		if flipcoin() {

			choose := rInt(0, len(level[a].roomrec))

			zobj := xobj{}

			zobj.cnt = level[a].roomrec[choose].cnt
			//zobj.solid = true
			zobj.fixed = true
			zobj.img = chestimg
			zobj.color = brightyellow()
			zobj.name = "chest"
			zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
			level[a].objs = append(level[a].objs, zobj)

		}

	}

	//level objs
	num := rInt(5, 15) //spike trap
	for {

		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()
		zobj.v1, _ = findcntr()

		zobj.fixed = true

		zobj.img = spikeimg
		zobj.color = randombluelight()
		zobj.fixed = true
		zobj.name = "spike trap"

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)

		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4

		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		zobj.cnt.X += tilesize

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)

		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4

		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(5, 15) //spring
	for {

		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()
		zobj.v1, _ = findcntr()

		zobj.fixed = true

		zobj.img = springimg
		zobj.color = rl.Magenta
		zobj.name = "spring"

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)

		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)
		num--
		if num == 0 {
			break
		}
	}
	num = rInt(5, 15) //weed
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.fixed = true
		chooseimg := rInt(1, 11)
		switch chooseimg {
		case 1:
			zobj.img = weed1img
		case 2:
			zobj.img = weed2img
		case 3:
			zobj.img = weed3img
		case 4:
			zobj.img = weed4img
		case 5:
			zobj.img = weed5img
		case 6:
			zobj.img = weed6img
		case 7:
			zobj.img = weed7img
		case 8:
			zobj.img = weed8img
		case 9:
			zobj.img = weed9img
		case 10:
			zobj.img = weed10img
		}

		zobj.color = randomgreen()
		zobj.name = "weed"

		zobj.usetype = rInt(1, 7)

		switch zobj.usetype {
		case 1:
			zobj.name2 = "antidote leaf"
		case 2:
			zobj.name2 = "resist poison bulb"
			zobj.amount = rInt(1, 11)
			zobj.amount = zobj.amount * 10
		case 3:
			zobj.name2 = "resist fire root"
			zobj.amount = rInt(1, 11)
			zobj.amount = zobj.amount * 10
		case 4:
			zobj.name2 = "cure disease seeds"
		case 5:
			zobj.name2 = "healing bark"
		case 6:
			zobj.name2 = "health boosting pollen"
		}

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(3, 8) //campfire
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.fixed = true
		zobj.img = campfireimg
		zobj.color = randomorange()
		zobj.name = "camp fire"

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(3, 6) //staircase
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.fixed = true
		chooseimg := rInt(1, 3)
		switch chooseimg {
		case 1:
			zobj.img = stairs1img
		case 2:
			zobj.img = stairs2img
		}

		zobj.color = randomcolor()
		zobj.name = "staircase"

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(5, 15) //empty potion
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.collect = true
		choose := rInt(0, len(potionemptyimgs))
		zobj.img = potionemptyimgs[choose]
		zobj.imgl = potionimgs[choose]
		zobj.color = randombluelight()
		zobj.name = "empty potion bottle"

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(10, 20) //quiver
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.img = quiverimg
		zobj.color = randomcolor()
		zobj.name = "quiver"
		zobj.kind = "ammo"
		zobj.numberof = rInt(50, 100)
		zobj.collect = true
		zobj.atk = 1

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(10, 20) //cactus
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.fixed = true
		zobj.img = cactusimg
		zobj.color = randomgreen()
		zobj.name = "cactus"

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(3, 8) //scroll
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.fixed = true
		zobj.collect = true
		chooseimg := rInt(1, 3)
		switch chooseimg {
		case 1:
			zobj.img = scroll1img
		case 2:
			zobj.img = scroll2img
		case 3:
			zobj.img = scroll3img
		}
		zobj.kind = "scroll"
		zobj.color = randomyellow()
		zobj.name = "scroll"

		zobj.scrollnum = rInt(0, 5)

		switch zobj.scrollnum {
		case 4:
			zobj.atk = rInt(2, 6)
			zobj.name2 = "poison gas scroll"
			zobj.uses = rInt(1, 5)
		case 3:
			zobj.atk = 4
			zobj.name2 = "meteors scroll"
			zobj.uses = rInt(1, 5)
		case 2:
			zobj.name2 = "identify scroll"
			zobj.uses = rInt(1, 5)
		case 1:
			zobj.name2 = "teleportation scroll"
			zobj.uses = rInt(1, 5)
		case 0:
			zobj.atk = rInt(2, 5)
			zobj.name2 = "ring of fire scroll"
			zobj.uses = rInt(1, 5)

		}

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(1, 4) //map
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.fixed = true
		zobj.collect = true
		chooseimg := rInt(1, 3)
		switch chooseimg {
		case 1:
			zobj.img = map1img
		case 2:
			zobj.img = map2img
		}
		zobj.color = randomyellow()
		zobj.name = "map"
		zobj.kind = "map"

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(5, 10) //plant
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.fixed = true
		zobj.img = plant1img
		zobj.color = randomgreen()
		zobj.name = "ornamental plant"

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(1, 4) //watering can
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.fixed = true
		zobj.img = wateringcanimg
		zobj.color = randombluelight()
		zobj.name = "magic watering can"

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(1, 4) //spade
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.fixed = true
		zobj.collect = true
		zobj.img = spadeimg
		zobj.imgl = spadeimgl
		zobj.origimg = spadeimg
		zobj.color = randombluelight()
		zobj.name = "spade"
		zobj.kind = "spade"

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(10, 20) //food
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.fixed = true
		zobj.collect = true
		chooseimg := rInt(1, 8)
		switch chooseimg {
		case 1:
			zobj.img = food1img
		case 2:
			zobj.img = food2img
		case 3:
			zobj.img = food3img
		case 4:
			zobj.img = food4img
		case 5:
			zobj.img = food5img
		case 6:
			zobj.img = food6img
		case 7:
			zobj.img = food7img
		}
		zobj.color = brightred()
		zobj.name = "food"

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(1, 4) //key
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.collect = true
		chooseimg := rInt(1, 4)
		switch chooseimg {
		case 1:
			zobj.img = key1img
		case 2:
			zobj.img = key2img
		case 3:
			zobj.img = key3img
		}
		zobj.color = randomcolor()
		zobj.name = "key"
		zobj.kind = "key"

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(1, 5) //cauldron
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.fixed = true
		zobj.img = cauldronimg
		zobj.color = brightyellow()
		zobj.name = "cauldron"

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(8, 12) //grave
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.fixed = true
		chooseimg := rInt(1, 5)
		switch chooseimg {
		case 1:
			zobj.img = grave1img
		case 2:
			zobj.img = grave2img
		case 3:
			zobj.img = grave3img
		case 4:
			zobj.img = grave4img
		}
		zobj.color = randomgrey()
		zobj.name = "grave"

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(8, 12) //bones
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.fixed = true
		chooseimg := rInt(1, 6)
		switch chooseimg {
		case 1:
			zobj.img = bones1img
		case 2:
			zobj.img = bones2img
		case 3:
			zobj.img = bones3img
		case 4:
			zobj.img = bones4img
		case 5:
			zobj.img = bones5img
		}
		zobj.color = randomgrey()
		zobj.name = "old bones"

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(1, 4) //bookcase
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.fixed = true
		chooseimg := rInt(1, 3)
		switch chooseimg {
		case 1:
			zobj.img = bookcase1img
		case 2:
			zobj.img = bookcase2img
		}
		zobj.color = randomcolor()
		zobj.name = "bookcase"

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(1, 5) //sign
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.fixed = true
		chooseimg := rInt(1, 10)
		switch chooseimg {
		case 1:
			zobj.img = sign1img
		case 2:
			zobj.img = sign2img
		case 3:
			zobj.img = sign3img
		case 4:
			zobj.img = sign4img
		case 5:
			zobj.img = sign5img
		case 6:
			zobj.img = sign6img
		case 7:
			zobj.img = sign7img
		case 8:
			zobj.img = sign8img
		case 9:
			zobj.img = sign9img
		}
		zobj.color = rl.Brown
		zobj.name = choosesignmsg()

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4
		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}
	num = rInt(5, 10) //torch
	for {
		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		zobj.fixed = true
		chooseimg := rInt(1, 4)
		switch chooseimg {
		case 1:
			zobj.img = torch1img
		case 2:
			zobj.img = torch2img
		case 3:
			zobj.img = torch3img
		}
		zobj.color = randomcolor()
		zobj.name = "torch"

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)

		zobj.boundrec = zobj.rec
		zobj.boundrec.X -= tilesize * 2
		zobj.boundrec.Y -= tilesize * 2
		zobj.boundrec.Width += tilesize * 4
		zobj.boundrec.Height += tilesize * 4

		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

		num--
		if num == 0 {
			break
		}
	}

	//water
	makewater()

}
func makesounds() { //MARK: makesounds

	campaud = rl.LoadSound("data/fx/campfire.mp3")
	stairsaud = rl.LoadSound("data/fx/downstairs.mp3")
	eataud = rl.LoadSound("data/fx/eat.mp3")
	bookcaseaud = rl.LoadSound("data/fx/bookcase.mp3")
	monsterhit1aud = rl.LoadSound("data/fx/monsterhit1.mp3")
	monsterhit2aud = rl.LoadSound("data/fx/monsterhit2.mp3")
	monsterhit3aud = rl.LoadSound("data/fx/monsterhit3.mp3")
	monsterhit4aud = rl.LoadSound("data/fx/monsterhit4.mp3")
	monsterhit5aud = rl.LoadSound("data/fx/monsterhit5.mp3")
	monsterhit6aud = rl.LoadSound("data/fx/monsterhit6.mp3")
	monsterhit7aud = rl.LoadSound("data/fx/monsterhit7.mp3")
	chestaud = rl.LoadSound("data/fx/openchest.mp3")
	weedaud = rl.LoadSound("data/fx/weed.mp3")

	digaud = rl.LoadSound("data/fx/spring.ogg")
	swingaud = rl.LoadSound("data/fx/swing_smack.ogg")
	zapaud = rl.LoadSound("data/fx/elec_loop_1.ogg")
	collectaud = rl.LoadSound("data/fx/collect.mp3")
	springaud = rl.LoadSound("data/fx/spring2.ogg")
	wateraud = rl.LoadSound("data/fx/watersplash.mp3")
	teleportpickupaud = rl.LoadSound("data/fx/teleportpickup.ogg")
	spikesaud = rl.LoadSound("data/fx/spikes.mp3")
	cactusaud = rl.LoadSound("data/fx/cactusdamage.mp3")
	playerdamageaud = rl.LoadSound("data/fx/playerdamage.mp3")
	gemaud = rl.LoadSound("data/fx/gem.ogg")

	intromusic = rl.LoadMusicStream("data/fx/crowd_noise_8.ogg")
	intromusic.Looping = true

	musicload := rl.LoadMusicStream("data/music/music1.ogg")
	backmusic = append(backmusic, musicload)
	musicload = rl.LoadMusicStream("data/music/music2.ogg")
	backmusic = append(backmusic, musicload)
	musicload = rl.LoadMusicStream("data/music/music3.ogg")
	backmusic = append(backmusic, musicload)
	musicload = rl.LoadMusicStream("data/music/music4.ogg")
	backmusic = append(backmusic, musicload)
	musicload = rl.LoadMusicStream("data/music/music5.ogg")
	backmusic = append(backmusic, musicload)
	musicload = rl.LoadMusicStream("data/music/music6.ogg")
	backmusic = append(backmusic, musicload)
	musicload = rl.LoadMusicStream("data/music/music7.ogg")
	backmusic = append(backmusic, musicload)
	musicload = rl.LoadMusicStream("data/music/music8.ogg")
	backmusic = append(backmusic, musicload)
	musicload = rl.LoadMusicStream("data/music/music9.ogg")
	backmusic = append(backmusic, musicload)
	musicload = rl.LoadMusicStream("data/music/music10.ogg")
	backmusic = append(backmusic, musicload)

	soundload := rl.LoadSound("data/fx/gabber_1.ogg")
	scrollaud = append(scrollaud, soundload)
	soundload = rl.LoadSound("data/fx/gabber_2.ogg")
	scrollaud = append(scrollaud, soundload)
	soundload = rl.LoadSound("data/fx/gabber_3.ogg")
	scrollaud = append(scrollaud, soundload)
	soundload = rl.LoadSound("data/fx/gabber_4.ogg")
	scrollaud = append(scrollaud, soundload)
	soundload = rl.LoadSound("data/fx/gabber_5.ogg")
	scrollaud = append(scrollaud, soundload)
	soundload = rl.LoadSound("data/fx/gabber_6.ogg")
	scrollaud = append(scrollaud, soundload)
	soundload = rl.LoadSound("data/fx/gabber_7.ogg")
	scrollaud = append(scrollaud, soundload)
	soundload = rl.LoadSound("data/fx/gabber_8.ogg")
	scrollaud = append(scrollaud, soundload)
	soundload = rl.LoadSound("data/fx/gabber_9.ogg")
	scrollaud = append(scrollaud, soundload)
	soundload = rl.LoadSound("data/fx/gabber_10.ogg")
	scrollaud = append(scrollaud, soundload)
	soundload = rl.LoadSound("data/fx/gabber_11.ogg")
	scrollaud = append(scrollaud, soundload)
	soundload = rl.LoadSound("data/fx/gabber_12.ogg")
	scrollaud = append(scrollaud, soundload)

}

func makelevel() { //MARK: makelevel

	checkgemskeys()

	cactusimmunity = false
	spikeimmunity = false
	killcount = 0
	level = nil
	fx = nil
	vismonsters = nil
	visroom = nil
	activmagic = nil
	activweapons = nil
	enemybullets = nil
	boss = nil
	activscroll = xobj{}
	questitemon = false
	questitemv2 = blankv2

	wallimg = wallimgs[rInt(0, len(wallimgs))]
	floorimg = floorimgs[rInt(0, len(floorimgs))]

	levmusic = backmusic[rInt(0, len(backmusic))]
	levmusic.Looping = true

	//player start room
	zroom := xroom{}
	zroom.num = 0
	zrec := xroomrec{}
	zrec.cnt = player.cnt

	multiw := float32(rInt(10, 20))
	multih := float32(rInt(10, 20))

	zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw*tilesize)/2), zrec.cnt.Y-((multih*tilesize)/2), (multiw * tilesize), (multih * tilesize))

	zrec.cnt = makereccnt(zrec.rec)
	zrec.collisrec = makecollisrec(zrec.rec)
	zroom.roomrec = append(zroom.roomrec, zrec)

	zroom.boundrec = makeboundrec(zroom)
	level = append(level, zroom)

	//make rooms
	num := 20

	for a := 1; a < num; a++ {

		zroom = xroom{}
		zroom.num = a
		zrec = xroomrec{}
		zrec.cnt = player.cnt

		multiw = float32(rInt(6, 14))
		multih = float32(rInt(6, 14))

		choose := rInt(1, 15)

		switch choose {

		case 14: //stepped right down

			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw*tilesize)/2), zrec.cnt.Y-((multih*tilesize)/2), (multiw * tilesize), (multih * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			num2 := rInt(2, 6)

			for {

				if int(multih)%2 != 0 {
					multih++
				}
				if int(multiw)%2 != 0 {
					multiw++
				}

				zrec.rec.X += (multiw / 2) * tilesize
				zrec.rec.Y += (multih / 2) * tilesize

				zrec.cnt = makereccnt(zrec.rec)
				zrec.collisrec = makecollisrec(zrec.rec)
				zroom.roomrec = append(zroom.roomrec, zrec)

				num2--
				if num2 == 0 {
					break
				}

			}

			zroom.boundrec = makeboundrec(zroom)
			zroom.xorig = zroom.boundrec.X
			zroom.yorig = zroom.boundrec.Y

		case 13: // pyramid down

			multih = 2
			multiw = float32(rInt(12, 19))

			if int(multiw)%2 == 0 {
				multiw += 1
			}

			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw*tilesize)/2), zrec.cnt.Y-((multih*tilesize)/2), (multiw * tilesize), (multih * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			for {
				zrec.rec.Y += multih * tilesize
				zrec.rec.X += tilesize
				zrec.rec.Width -= tilesize * 2

				zrec.cnt = makereccnt(zrec.rec)
				zrec.collisrec = makecollisrec(zrec.rec)
				zroom.roomrec = append(zroom.roomrec, zrec)

				if zrec.rec.Width <= tilesize*4 {
					break
				}

			}
			zroom.boundrec = makeboundrec(zroom)
			zroom.xorig = zroom.boundrec.X
			zroom.yorig = zroom.boundrec.Y

		case 12: // pyramid up

			multih = 2
			multiw = float32(rInt(12, 19))

			if int(multiw)%2 == 0 {
				multiw += 1
			}

			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw*tilesize)/2), zrec.cnt.Y-((multih*tilesize)/2), (multiw * tilesize), (multih * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			for {
				zrec.rec.Y -= multih * tilesize
				zrec.rec.X += tilesize
				zrec.rec.Width -= tilesize * 2

				zrec.cnt = makereccnt(zrec.rec)
				zrec.collisrec = makecollisrec(zrec.rec)
				zroom.roomrec = append(zroom.roomrec, zrec)

				if zrec.rec.Width <= tilesize*4 {
					break
				}

			}

			zroom.boundrec = makeboundrec(zroom)
			zroom.xorig = zroom.boundrec.X
			zroom.yorig = zroom.boundrec.Y

		case 11: //vertical passage of rooms
			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw*tilesize)/2), zrec.cnt.Y-((multih*tilesize)/2), (multiw * tilesize), (multih * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			y1 := zrec.cnt.Y
			y2 := zrec.cnt.Y

			num2 := rInt(2, 6)

			for {
				zrec.rec.Y -= multih * tilesize

				multih2 := multih

				if int(multih2)%2 != 0 {
					multih2++
				}
				zrec.rec.Y -= (multih2 / 2) * tilesize

				zrec.cnt = makereccnt(zrec.rec)
				zrec.collisrec = makecollisrec(zrec.rec)
				zroom.roomrec = append(zroom.roomrec, zrec)
				y2 = zrec.cnt.Y
				num2--
				if num2 == 0 {
					break
				}
			}

			multiw = 4

			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw*tilesize)/2), y2, multiw*tilesize, y1-y2)

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zroom.boundrec = makeboundrec(zroom)
			zroom.xorig = zroom.boundrec.X
			zroom.yorig = zroom.boundrec.Y

		case 10: //horizontal passage of rooms
			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw*tilesize)/2), zrec.cnt.Y-((multih*tilesize)/2), (multiw * tilesize), (multih * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			x1 := zrec.cnt.X
			x2 := zrec.cnt.X

			num2 := rInt(2, 6)

			for {
				zrec.rec.X += multiw * tilesize

				multiw2 := multiw

				if int(multiw2)%2 != 0 {
					multiw2++
				}
				zrec.rec.X += (multiw2 / 2) * tilesize

				zrec.cnt = makereccnt(zrec.rec)
				zrec.collisrec = makecollisrec(zrec.rec)
				zroom.roomrec = append(zroom.roomrec, zrec)
				x2 = zrec.cnt.X

				num2--
				if num2 == 0 {
					break
				}
			}

			multih = 4

			zrec.rec = rl.NewRectangle(x1, zrec.cnt.Y-((multih*tilesize)/2), x2-x1, (multih * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zroom.boundrec = makeboundrec(zroom)
			zroom.xorig = zroom.boundrec.X
			zroom.yorig = zroom.boundrec.Y

		case 9: //4 recs with side & center passages
			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw*tilesize)/2), zrec.cnt.Y-((multih*tilesize)/2), (multiw * tilesize), (multih * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.X += (multiw * tilesize) * 2

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.Y += (multih * tilesize) * 2

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.X -= (multiw * tilesize) * 2

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			multiw2 := float32(4)
			multih2 := (multih * tilesize) * 2

			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw2*tilesize)/2), zrec.cnt.Y-multih2, (multiw2 * tilesize), (multih2))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.X += (multiw * tilesize) * 2

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.X -= (multiw * tilesize) * 2
			zrec.rec.X += (multiw2 * tilesize) / 2
			zrec.rec.Y -= (multiw2 * tilesize) / 2

			zrec.rec.Width = (multiw * tilesize) * 2
			zrec.rec.Height = multiw2 * tilesize

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.Y += (multih * tilesize) * 2

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.Y -= (multih * tilesize)

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zroom.boundrec = makeboundrec(zroom)
			zroom.xorig = zroom.boundrec.X
			zroom.yorig = zroom.boundrec.Y

		case 8: //4 recs with center passage
			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw*tilesize)/2), zrec.cnt.Y-((multih*tilesize)/2), (multiw * tilesize), (multih * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.X += (multiw * tilesize) * 2

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.Y += (multih * tilesize) * 2

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.X -= (multiw * tilesize) * 2

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			multiw2 := float32(4)
			multih2 := (multih * tilesize) * 2

			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw2*tilesize)/2), zrec.cnt.Y-multih2, (multiw2 * tilesize), (multih2))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.X += (multiw * tilesize) * 2

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.X -= (multiw * tilesize) * 2
			zrec.rec.X += (multiw2 * tilesize) / 2
			zrec.rec.Y -= (multiw2 * tilesize) / 2

			zrec.rec.Width = (multiw * tilesize) * 2
			zrec.rec.Height = multiw2 * tilesize

			zrec.rec.Y += (multih * tilesize)

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zroom.boundrec = makeboundrec(zroom)
			zroom.xorig = zroom.boundrec.X
			zroom.yorig = zroom.boundrec.Y

		case 7: //4 recs with side passages
			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw*tilesize)/2), zrec.cnt.Y-((multih*tilesize)/2), (multiw * tilesize), (multih * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.X += (multiw * tilesize) * 2

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.Y += (multih * tilesize) * 2

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.X -= (multiw * tilesize) * 2

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			multiw2 := float32(4)
			multih2 := (multih * tilesize) * 2

			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw2*tilesize)/2), zrec.cnt.Y-multih2, (multiw2 * tilesize), (multih2))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.X += (multiw * tilesize) * 2

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.X -= (multiw * tilesize) * 2
			zrec.rec.X += (multiw2 * tilesize) / 2
			zrec.rec.Y -= (multiw2 * tilesize) / 2

			zrec.rec.Width = (multiw * tilesize) * 2
			zrec.rec.Height = multiw2 * tilesize

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.Y += (multih * tilesize) * 2

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zroom.boundrec = makeboundrec(zroom)
			zroom.xorig = zroom.boundrec.X
			zroom.yorig = zroom.boundrec.Y

		case 6: //stepped cross

			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw*tilesize)/2), zrec.cnt.Y-((multih*tilesize)/2), (multiw * tilesize), (multih * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			num2 := 2

			for {
				if int(multih)%2 != 0 {
					multih++
				}
				if int(multiw)%2 != 0 {
					multiw++
				}
				zrec.rec.X += (multiw / 2) * tilesize
				zrec.rec.Y -= (multih / 2) * tilesize

				zrec.cnt = makereccnt(zrec.rec)
				zrec.collisrec = makecollisrec(zrec.rec)
				zroom.roomrec = append(zroom.roomrec, zrec)

				num2--
				if num2 == 0 {
					break
				}
			}
			num2 = 2
			for {
				if int(multih)%2 != 0 {
					multih++
				}
				if int(multiw)%2 != 0 {
					multiw++
				}
				zrec.rec.X += (multiw / 2) * tilesize
				zrec.rec.Y += (multih / 2) * tilesize

				zrec.cnt = makereccnt(zrec.rec)
				zrec.collisrec = makecollisrec(zrec.rec)
				zroom.roomrec = append(zroom.roomrec, zrec)

				num2--
				if num2 == 0 {
					break
				}
			}

			num2 = 2
			for {
				if int(multih)%2 != 0 {
					multih++
				}
				if int(multiw)%2 != 0 {
					multiw++
				}
				zrec.rec.X -= (multiw / 2) * tilesize
				zrec.rec.Y += (multih / 2) * tilesize

				zrec.cnt = makereccnt(zrec.rec)
				zrec.collisrec = makecollisrec(zrec.rec)
				zroom.roomrec = append(zroom.roomrec, zrec)

				num2--
				if num2 == 0 {
					break
				}
			}
			num2 = 2
			for {
				if int(multih)%2 != 0 {
					multih++
				}
				if int(multiw)%2 != 0 {
					multiw++
				}
				zrec.rec.X -= (multiw / 2) * tilesize
				zrec.rec.Y -= (multih / 2) * tilesize

				zrec.cnt = makereccnt(zrec.rec)
				zrec.collisrec = makecollisrec(zrec.rec)
				zroom.roomrec = append(zroom.roomrec, zrec)

				num2--
				if num2 == 0 {
					break
				}
			}
			zroom.boundrec = makeboundrec(zroom)
			zroom.xorig = zroom.boundrec.X
			zroom.yorig = zroom.boundrec.Y
		case 5: //stepped right up

			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw*tilesize)/2), zrec.cnt.Y-((multih*tilesize)/2), (multiw * tilesize), (multih * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			num2 := rInt(2, 6)

			for {

				if int(multih)%2 != 0 {
					multih++
				}
				if int(multiw)%2 != 0 {
					multiw++
				}

				zrec.rec.X += (multiw / 2) * tilesize
				zrec.rec.Y -= (multih / 2) * tilesize

				zrec.cnt = makereccnt(zrec.rec)
				zrec.collisrec = makecollisrec(zrec.rec)
				zroom.roomrec = append(zroom.roomrec, zrec)

				num2--
				if num2 == 0 {
					break
				}

			}

			zroom.boundrec = makeboundrec(zroom)
			zroom.xorig = zroom.boundrec.X
			zroom.yorig = zroom.boundrec.Y

		case 4: //cross
			multiw = float32(rInt(4, 6))
			multih = float32(rInt(10, 15))

			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw*tilesize)/2), zrec.cnt.Y-((multih*tilesize)/2), (multiw * tilesize), (multih * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multih*tilesize)/2), zrec.cnt.Y-((multiw*tilesize)/2), (multih * tilesize), (multiw * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zroom.boundrec = makeboundrec(zroom)
			zroom.xorig = zroom.boundrec.X
			zroom.yorig = zroom.boundrec.Y

		case 3: //rec center 4 recs mid border
			if int(multih)%2 != 0 {
				multih++
			}
			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw*tilesize)/2), zrec.cnt.Y-((multih*tilesize)/2), (multiw * tilesize), (multih * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			cntorig := zrec.cnt

			multiw -= 2
			multih -= 2

			zrec.rec = rl.NewRectangle(cntorig.X-((multiw*tilesize)/2), zrec.rec.Y-((multih*tilesize)/2), (multiw * tilesize), (multih * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.Y += (multih + 2) * tilesize

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec = rl.NewRectangle((cntorig.X-(((multiw+2)*tilesize)/2))-((multih/2)*tilesize), cntorig.Y-((multiw*tilesize)/2), (multih * tilesize), (multiw * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.X += (multiw + 2) * tilesize

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zroom.boundrec = makeboundrec(zroom)
			zroom.xorig = zroom.boundrec.X
			zroom.yorig = zroom.boundrec.Y

		case 2: //rec center 4 recs corners
			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw*tilesize)/2), zrec.cnt.Y-((multih*tilesize)/2), (multiw * tilesize), (multih * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			multiw -= 2
			multih -= 2

			zrec.rec = rl.NewRectangle(zrec.rec.X-((multiw*tilesize)/2), zrec.rec.Y-((multih*tilesize)/2), (multiw * tilesize), (multih * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.X += (multiw + 2) * tilesize

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.Y += (multih + 2) * tilesize

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec.X -= (multiw + 2) * tilesize

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)

			zroom.boundrec = makeboundrec(zroom)
			zroom.xorig = zroom.boundrec.X
			zroom.yorig = zroom.boundrec.Y

		case 1: //rec
			zrec.rec = rl.NewRectangle(zrec.cnt.X-((multiw*tilesize)/2), zrec.cnt.Y-((multih*tilesize)/2), (multiw * tilesize), (multih * tilesize))

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)
			zroom.roomrec = append(zroom.roomrec, zrec)
			zroom.boundrec = makeboundrec(zroom)
			zroom.xorig = zroom.boundrec.X
			zroom.yorig = zroom.boundrec.Y
		}

		level = append(level, zroom)

	}

	//place rooms

	roomnum := 1
	countbreak := 1000
	for {

		addroom := false

		chooseroom := rInt(0, roomnum)

		checkrec := level[roomnum].boundrec
		checkrec.X = level[chooseroom].boundrec.X
		checkrec.Y = level[chooseroom].boundrec.Y

		side := rInt(1, 5)

		switch side {
		case 1: //top

			checkrec.Y -= level[roomnum].boundrec.Height

			multi1 := int((level[chooseroom].boundrec.Width / tilesize))

			multi1 = rInt(-multi1, multi1+1)

			checkrec.X += float32(multi1) * tilesize

			collides := false
			for a := 0; a < roomnum; a++ {

				if rl.CheckCollisionRecs(checkrec, level[a].boundrec) {
					collides = true
				}

			}

			if !collides {
				addroom = true
			}

		case 3: //bottom

			checkrec.Y += level[chooseroom].boundrec.Height

			multi1 := int((level[chooseroom].boundrec.Width / tilesize))

			multi1 = rInt(-multi1, multi1+1)

			checkrec.X += float32(multi1) * tilesize

			collides := false
			for a := 0; a < roomnum; a++ {

				if rl.CheckCollisionRecs(checkrec, level[a].boundrec) {
					collides = true

				}
			}

			if !collides {
				addroom = true
			}

		case 4: //left

			checkrec.X -= level[roomnum].boundrec.Width

			multi1 := int((level[chooseroom].boundrec.Height / tilesize))

			multi1 = rInt(-multi1, multi1+1)

			checkrec.Y += float32(multi1) * tilesize

			collides := false
			for a := 0; a < roomnum; a++ {

				if rl.CheckCollisionRecs(checkrec, level[a].boundrec) {
					collides = true

				}
			}

			if !collides {
				addroom = true
			}

		case 2: //right

			checkrec.X += level[chooseroom].boundrec.Width

			multi1 := int((level[chooseroom].boundrec.Height / tilesize))

			multi1 = rInt(-multi1, multi1+1)

			checkrec.Y += float32(multi1) * tilesize

			collides := false
			for a := 0; a < roomnum; a++ {

				if rl.CheckCollisionRecs(checkrec, level[a].boundrec) {
					collides = true
				}

			}

			if !collides {
				addroom = true
			}

		}

		if addroom {
			moveroom(roomnum, checkrec.X, checkrec.Y)
			roomnum++
		}

		if roomnum == len(level) {
			break
		}

		countbreak--

		if countbreak == 0 {
			break
		}

	}

	origlevellen = len(level)

	//make passages
	for a := 1; a < origlevellen; a++ {
		makepassage(a)
	}

	//find xy width length level

	xl := level[0].boundrec.X
	yt := level[0].boundrec.Y

	xr := level[0].boundrec.X + level[0].boundrec.Width
	yb := level[0].boundrec.Y + level[0].boundrec.Height

	for a := 1; a < len(level); a++ {

		if level[a].boundrec.X < xl {
			xl = level[a].boundrec.X
		}
		if level[a].boundrec.X+level[a].boundrec.Width > xr {
			xr = level[a].boundrec.X + level[a].boundrec.Width
		}

		if level[a].boundrec.Y < yt {
			yt = level[a].boundrec.Y
		}
		if level[a].boundrec.Y+level[a].boundrec.Height > yb {
			yb = level[a].boundrec.Y + level[a].boundrec.Height
		}
	}

	levtl = rl.NewVector2(xl, yt)
	levtr = rl.NewVector2(xr, yt)

	levbl = rl.NewVector2(xl, yb)
	levbr = rl.NewVector2(xr, yb)

	levwid = xr - xl
	levheig = yb - yt

	levboundrec = rl.NewRectangle(xl, yt, levwid, levheig)
	backgrec = levboundrec
	backgrec.X -= tilesize * 2
	backgrec.Y -= tilesize * 2
	backgrec.Width += tilesize * 4
	backgrec.Height += tilesize * 4

	currentlevelnum++
	maxlevelreached = currentlevelnum
	score += currentlevelnum * 100

	//make inner recs
	makeinnerrecs()
	//make objs
	makestartobjs()
	makebackobjs()
	//make monsters
	makemonsters()

	cleanlevel()

	selpoint = player.cnt
	selrec = rl.NewRectangle(selpoint.X-tilesize/4, selpoint.Y-tilesize/4, tilesize/2, tilesize/2)

	player.cnt = level[0].roomrec[0].cnt

	newmsg("you have arrived in dungeon LEVEL " + fmt.Sprint(currentlevelnum) + " you cannot return to the previous level")
	refillcost = currentlevelnum

	pause = false
}

func makeimgs() { //MARK: makeimgs

	x := float32(485)
	y := float32(398)
	for {
		mushroombossl = append(mushroombossl, rl.NewRectangle(x, y, 32, 32))
		x -= 32
		if x < 2 {
			break
		}
	}

	x = float32(2)
	y = float32(445)
	for {
		mushroomboss = append(mushroomboss, rl.NewRectangle(x, y, 32, 32))
		x += 32
		if x > 484 {
			break
		}
	}

	x = float32(196)
	y = float32(495)
	for {
		radishboss = append(radishboss, rl.NewRectangle(x, y, 38, 38))
		x += 38
		if x > 396 {
			break
		}
	}

	x = float32(391)
	y = float32(555)
	for {
		radishbossl = append(radishbossl, rl.NewRectangle(x, y, 38, 38))
		x -= 38
		if x < 190 {
			break
		}
	}

	x = float32(25)
	y = float32(720)
	for {
		spikeboss = append(spikeboss, rl.NewRectangle(x, y, 44, 44))
		x += 44
		if x > 344 {
			break
		}
	}

	x = float32(332)
	y = float32(782)
	for {
		spikebossl = append(spikebossl, rl.NewRectangle(x, y, 44, 44))
		x -= 44
		if x < 12 {
			break
		}
	}

	x = float32(1023)
	y = float32(542)
	for {
		skullboss = append(skullboss, rl.NewRectangle(x, y, 52, 52))
		x += 52
		if x > 1396 {
			break
		}
	}

	x = float32(1391)
	y = float32(598)
	for {
		skullbossl = append(skullbossl, rl.NewRectangle(x, y, 52, 52))
		x -= 52
		if x < 1020 {
			break
		}
	}

	x = float32(552)
	y = float32(448)
	for {
		ghostboss = append(ghostboss, rl.NewRectangle(x, y, 44, 44))
		x += 44
		if x > 956 {
			break
		}
	}

	x = float32(950)
	y = float32(492)
	for {
		ghostbossl = append(ghostbossl, rl.NewRectangle(x, y, 44, 44))
		x -= 44
		if x < 536 {
			break
		}
	}

	x = float32(727)
	y = float32(394)
	for {
		reaperboss = append(reaperboss, rl.NewRectangle(x, y, 48, 48))
		x += 48
		if x > 934 {
			break
		}
	}

	x = float32(924)
	y = float32(342)
	for {
		reaperbossl = append(reaperbossl, rl.NewRectangle(x, y, 48, 48))
		x -= 48
		if x < 724 {
			break
		}
	}

	x = float32(142)
	y = float32(611)
	for {
		orcboss = append(orcboss, rl.NewRectangle(x, y, 24, 24))
		x += 24
		if x > 290 {
			break
		}
	}

	x = float32(285)
	y = float32(653)
	for {
		orcbossl = append(orcbossl, rl.NewRectangle(x, y, 24, 24))
		x -= 24
		if x < 138 {
			break
		}
	}

	x = float32(553)
	y = float32(543)
	for {
		slimeboss = append(slimeboss, rl.NewRectangle(x, y, 44, 44))
		x += 44
		if x > 960 {
			break
		}
	}

	x = float32(953)
	y = float32(602)
	for {
		slimebossl = append(slimebossl, rl.NewRectangle(x, y, 44, 44))
		x -= 44
		if x < 550 {
			break
		}
	}

	x = float32(291)
	y = float32(329)
	for {
		dinoboss = append(dinoboss, rl.NewRectangle(x, y, 24, 24))
		x += 24
		if x > 605 {
			break
		}
	}

	x = float32(603)
	y = float32(360)
	for {
		dinobossl = append(dinobossl, rl.NewRectangle(x, y, 24, 24))
		x -= 24
		if x < 288 {
			break
		}
	}

	x = float32(0)
	y = float32(0)

	for {
		wallimgs = append(wallimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x >= 480 {
			break
		}
	}

	x = float32(1080)
	y = float32(199)
	for {
		helmetimgs = append(helmetimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 1176 {
			break
		}
	}

	x = float32(1080)
	y = float32(215)
	for {
		vestimgs = append(vestimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 1144 {
			break
		}
	}

	x = float32(1160)
	y = float32(215)
	for {
		robeimgs = append(robeimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 1176 {
			break
		}
	}

	x = float32(1260)
	y = float32(200)
	for {
		crownimgs = append(crownimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 1276 {
			break
		}
	}

	x = float32(1192)
	y = float32(199)
	for {
		bootimgs = append(bootimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 1208 {
			x = 1192
			y += 16
		}
		if y > 215 {
			break
		}
	}

	x = float32(1224)
	y = float32(199)
	for {
		gloveimgs = append(gloveimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 1240 {
			x = 1224
			y += 16
		}
		if y > 215 {
			break
		}
	}

	x = float32(810)
	y = float32(250)

	for {
		gemimgs = append(gemimgs, rl.NewRectangle(x, y, 18, 18))
		x += 24
		if x > 858 {
			y += 24
			x = 810
		}
		if y > 298 {
			break
		}
	}

	x = float32(890)
	y = float32(251)

	for {
		jewelryimgs = append(jewelryimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 938 {
			y += 16
			x = 890
		}
		if y > 283 {
			break
		}
	}

	x = float32(592)
	y = float32(247)

	for {
		potionimgs = append(potionimgs, rl.NewRectangle(x, y, 18, 18))
		x += 24
		if x > 640 {
			y += 24
			x = 592
		}
		if y > 295 {
			break
		}
	}

	x = float32(668)
	y = float32(247)

	for {
		potionemptyimgs = append(potionemptyimgs, rl.NewRectangle(x, y, 18, 18))
		x += 24
		if x > 716 {
			y += 24
			x = 668
		}
		if y > 295 {
			break
		}
	}

	x = float32(0)
	y = float32(16)

	for {
		floorimgs = append(floorimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x >= 720 {
			break
		}
	}

	x = float32(20)
	y = float32(186)
	for {
		swordimgs = append(swordimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 244 {
			break
		}
	}

	x = float32(20)
	y = float32(203)
	for {
		swordimgsl = append(swordimgsl, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 244 {
			break
		}
	}

	x = float32(261)
	y = float32(186)
	for {
		daggerimgs = append(daggerimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 328 {
			break
		}
	}

	x = float32(261)
	y = float32(202)
	for {
		daggerimgsl = append(daggerimgsl, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 328 {
			break
		}
	}

	x = float32(342)
	y = float32(190)
	for {
		clubimgs = append(clubimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 390 {
			break
		}
	}

	x = float32(342)
	y = float32(206)
	for {
		clubimgsl = append(clubimgsl, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 390 {
			break
		}
	}

	x = float32(407)
	y = float32(190)
	for {
		scytheimgs = append(scytheimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 423 {
			break
		}
	}

	x = float32(406)
	y = float32(206)
	for {
		scytheimgsl = append(scytheimgsl, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 423 {
			break
		}
	}

	x = float32(440)
	y = float32(190)
	for {
		maceimgs = append(maceimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 537 {
			break
		}
	}

	x = float32(440)
	y = float32(207)
	for {
		maceimgsl = append(maceimgsl, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 537 {
			break
		}
	}

	x = float32(553)
	y = float32(190)
	for {
		wandimgs = append(wandimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 649 {
			break
		}
	}

	x = float32(553)
	y = float32(206)
	for {
		wandimgsl = append(wandimgsl, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 649 {
			break
		}
	}

	x = float32(665)
	y = float32(190)
	for {
		axeimgs = append(axeimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 795 {
			break
		}
	}

	x = float32(665)
	y = float32(206)
	for {
		axeimgsl = append(axeimgsl, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 795 {
			break
		}
	}

	x = float32(811)
	y = float32(192)
	for {
		throwingaxeimgs = append(throwingaxeimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 827 {
			break
		}
	}

	x = float32(811)
	y = float32(206)
	for {
		throwingaxeimgsl = append(throwingaxeimgsl, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 827 {
			break
		}
	}

	x = float32(1021)
	y = float32(192)
	for {
		spearimgs = append(spearimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 1037 {
			break
		}
	}

	x = float32(1021)
	y = float32(207)
	for {
		spearimgsl = append(spearimgsl, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 1037 {
			break
		}
	}

	x = float32(872)
	y = float32(190)
	for {
		crossbowimgs = append(crossbowimgs, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 904 {
			break
		}
	}

	x = float32(873)
	y = float32(206)
	for {
		crossbowimgsl = append(crossbowimgsl, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x > 905 {
			break
		}
	}

	x = float32(972)
	y = float32(253)
	for {
		bowimgs = append(bowimgs, rl.NewRectangle(x, y, 16, 16))
		x += 17
		if x > 1040 {
			break
		}
	}

	x = float32(972)
	y = float32(270)
	for {
		bowimgsl = append(bowimgsl, rl.NewRectangle(x, y, 16, 16))
		x += 17
		if x > 1040 {
			break
		}
	}

	x = float32(630)
	y = float32(841)
	for {
		emoteimgs = append(emoteimgs, rl.NewRectangle(x, y, 32, 32))
		x += 96
		if x > 1562 {
			x = 630
			y += 32
		}
		if y > 1168 {
			break
		}
	}

}
func makelegendaryitem() { //MARK: makelegendaryitem

	zobj := xobj{}
	roomnumobjs := 0
	zobj.cnt, roomnumobjs = findcntr()
	questitemroomnum = roomnumobjs

	questitemv2 = zobj.cnt

	zobj.kind = "armor"
	zobj.collect = true
	zobj.questitem = true
	zobj.legendary = true
	zobj.name = "legendary armor"
	zobj.color = brightyellow()
	zobj.armorsetnum = rInt(1, 7)
	zobj.ability = rInt(1, 9)

	switch zobj.ability {
	case 1:
		zobj.name4 = "health regen"
	case 2:
		zobj.name4 = "vampirism"
	case 3:
		zobj.name4 = "thorns"
	case 4:
		zobj.name4 = "speed"
	case 5:
		zobj.name4 = "fire trail"
	case 6:
		zobj.name4 = "teleport"
		zobj.timer = rInt32(int(fps*10), int(fps*20))
	case 7:
		zobj.name4 = "rainbow"
	case 8:
		zobj.name4 = "identify"

	}

	choose := rInt(1, 7)
	switch choose {
	case 1:
		zobj.img = helmetimgs[rInt(0, len(helmetimgs))]
		zobj.usetype = 1

		zobj.name2 = "legendary helmet of " + zobj.name4 + " - part of armor set " + (fmt.Sprint(zobj.armorsetnum))
		zobj.name3 = "helmet"
	case 2:
		zobj.img = bootimgs[rInt(0, len(bootimgs))]
		zobj.usetype = 2

		zobj.name2 = "legendary boots of " + zobj.name4 + " - part of armor set " + (fmt.Sprint(zobj.armorsetnum))
		zobj.name3 = "boots"
	case 3:
		zobj.img = gloveimgs[rInt(0, len(gloveimgs))]
		zobj.usetype = 3

		zobj.name2 = "legendary gloves of " + zobj.name4 + " - part of armor set " + (fmt.Sprint(zobj.armorsetnum))
		zobj.name3 = "gloves"
	case 4:
		zobj.img = vestimgs[rInt(0, len(vestimgs))]
		zobj.usetype = 4

		zobj.name2 = "legendary vest of " + zobj.name4 + " - part of armor set " + (fmt.Sprint(zobj.armorsetnum))
		zobj.name3 = "vest"
	case 5:
		zobj.img = robeimgs[rInt(0, len(robeimgs))]
		zobj.usetype = 4

		zobj.name2 = "legendary robe of " + zobj.name4 + " - part of armor set " + (fmt.Sprint(zobj.armorsetnum))
		zobj.name3 = "vest"
	case 6:
		zobj.img = crownimgs[rInt(0, len(crownimgs))]
		zobj.usetype = 1

		zobj.name2 = "legendary crown of " + zobj.name4 + " - part of armor set " + (fmt.Sprint(zobj.armorsetnum))
		zobj.name3 = "helmet"
	}

	zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)
	zobj.boundrec = zobj.rec
	zobj.boundrec.X -= tilesize * 2
	zobj.boundrec.Y -= tilesize * 2
	zobj.boundrec.Width += tilesize * 4
	zobj.boundrec.Height += tilesize * 4
	level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)

	questitemon = true
}
func makeblood() { //MARK: makeblood

	num := rInt(30, 50)

	for {

		zcirc := xcircle{}

		zcirc.v2 = rl.NewVector2(rFloat32(0, scrwf32), rFloat32(0, scrhf32))
		zcirc.rad = rFloat32(40, 120)
		zcirc.color = brightred()
		zcirc.fade = 1.0

		blood = append(blood, zcirc)

		num--
		if num == 0 {
			break

		}
	}

}
func makerangerec(weapon xobj) rl.Rectangle { //MARK: makerangerec
	rec := rl.NewRectangle(0, 0, weapon.meleerange*2, weapon.meleerange*2)
	return rec
}
func makepet() { //MARK: makepet

	zpet := xpet{}
	zpet.cnt = player.cnt
	zpet.cnt.X += tilesize
	zpet.timer = rInt(60, 120)
	zpet.rec = rl.NewRectangle(zpet.cnt.X-tilesize/2, zpet.cnt.Y-tilesize/2, tilesize, tilesize)
	zpet.vel = player.vel
	zpet.name = petnames[rInt(0, len(petnames))]

	choose := rInt(1, 6)

	switch choose {

	case 5:
		zpet.imgi = dog2iimg
		zpet.imgr = dog2rimg
		zpet.imgl = dog2limg
		zpet.frames = 3
		zpet.startxr = zpet.imgr.X
		zpet.startxl = zpet.imgl.X
		zpet.endxr = zpet.startxr + (zpet.imgr.Width * zpet.frames)
		zpet.endxl = zpet.startxl - (zpet.imgr.Width * zpet.frames)

	case 4:
		zpet.imgi = dog1iimg
		zpet.imgr = dog1rimg
		zpet.imgl = dog1limg
		zpet.frames = 3
		zpet.startxr = zpet.imgr.X
		zpet.startxl = zpet.imgl.X
		zpet.endxr = zpet.startxr + (zpet.imgr.Width * zpet.frames)
		zpet.endxl = zpet.startxl - (zpet.imgr.Width * zpet.frames)

	case 3:
		zpet.imgi = sheepiimg
		zpet.imgr = sheeprimg
		zpet.imgl = sheeplimg
		zpet.frames = 3
		zpet.startxr = zpet.imgr.X
		zpet.startxl = zpet.imgl.X
		zpet.endxr = zpet.startxr + (zpet.imgr.Width * zpet.frames)
		zpet.endxl = zpet.startxl - (zpet.imgr.Width * zpet.frames)

	case 2:
		zpet.imgi = mouse2iimg
		zpet.imgr = mouse2rimg
		zpet.imgl = mouse2limg
		zpet.frames = 3
		zpet.startxr = zpet.imgr.X
		zpet.startxl = zpet.imgl.X
		zpet.endxr = zpet.startxr + (zpet.imgr.Width * zpet.frames)
		zpet.endxl = zpet.startxl - (zpet.imgr.Width * zpet.frames)

	case 1:
		zpet.imgi = mouse1iimg
		zpet.imgr = mouse1rimg
		zpet.imgl = mouse1limg
		zpet.frames = 3
		zpet.startxr = zpet.imgr.X
		zpet.startxl = zpet.imgl.X
		zpet.endxr = zpet.startxr + (zpet.imgr.Width * zpet.frames)
		zpet.endxl = zpet.startxl - (zpet.imgr.Width * zpet.frames)

	}

	pets = append(pets, zpet)

}
func makelaser() { //MARK: makelaser

	zobj := xobj{}
	zobj.noimg = true
	zobj.usetype = 1

	zobj.v1 = player.cnt

	zobj.dirx, zobj.diry = weapondir(8)
	zobj.v2.X = zobj.v1.X + (zobj.dirx * 5)
	zobj.v2.Y = zobj.v1.Y + (zobj.diry * 5)

	zobj.timer = fps * 6

	activweapons = append(activweapons, zobj)

}
func makeorbit() { //MARK: makeorbit

	zobj := xobj{}
	zobj.noimg = true
	zobj.usetype = 2

	zobj.v1 = player.cnt
	zobj.v2 = player.cnt

	zobj.v1.Y -= tilesize * 3

	zobj.angle = angle2points(zobj.v2, zobj.v1)

	zobj.timer = fps * 10

	activweapons = append(activweapons, zobj)

}
func makefrogs() { //MARK: makefrogs

	num := rInt(10, 20)

	for {

		zobj := xobj{}
		zobj.usetype = 3

		zobj.img = frogimg

		zobj.timer = fps * 10

		zobj.color = rl.White

		zobj.cnt = finrandpointborderrec()
		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)

		activweapons = append(activweapons, zobj)

		num--
		if num == 0 {
			break
		}
	}

	frograintimer = fps * 3
}
func makelightning() bool {

	decrease := false

	var enemycnt []rl.Vector2

	for a := 0; a < len(vismonsters); a++ {
		if rl.CheckCollisionPointRec(vismonsters[a].cnt, borderrec) {
			enemycnt = append(enemycnt, vismonsters[a].cnt)

		}
	}

	if len(enemycnt) > 0 {
		decrease = true
		zobj := xobj{}
		zobj.noimg = true
		zobj.usetype = 4

		zobj.v1 = player.cnt

		zobj.timer = fps * 6

		activweapons = append(activweapons, zobj)

		for a := 0; a < len(enemycnt); a++ {

			zobj := xobj{}
			zobj.noimg = true
			zobj.usetype = 4

			zobj.v1 = enemycnt[a]

			zobj.timer = fps * 6

			activweapons = append(activweapons, zobj)
			if soundfxon && !mute {
				rl.PlaySoundMulti(zapaud)
			}

		}

	} else {
		newmsg("there are no monsters on screen to zap...")
	}

	return decrease
}
func makepotions() { //MARK: makepotions

	var weeds []xobj
	var bottles []xobj
	var invennumbottles []int
	var invennumweeds []int

	for a := 0; a < len(inven); a++ {

		if inven[a].name == "empty potion bottle" {
			bottles = append(bottles, inven[a])
			invennumbottles = append(invennumbottles, a)
		}
		if inven[a].name == "weed" {
			weeds = append(weeds, inven[a])
			invennumweeds = append(invennumweeds, a)
		}
	}

	if len(bottles) > 0 {
		if len(weeds) > 0 {
			if len(bottles) <= len(weeds) {
				for a := 0; a < len(bottles); a++ {
					zobj := xobj{}
					zobj.img = bottles[a].imgl
					zobj.name = "potion"
					zobj.color = randomcolor()
					zobj.rec = inven[invennumbottles[a]].rec
					zobj.cnt = inven[invennumbottles[a]].cnt
					zobj.kind = "potion"
					switch inven[invennumweeds[a]].name2 {
					case "health boosting pollen":
						zobj.name2 = "hp max+ potion"
						zobj.color = randomred()
					case "healing bark":
						zobj.name2 = "healing potion"
						zobj.color = randomred()
					case "cure disease seeds":
						zobj.name2 = "cure disease potion"
						zobj.color = randomyellow()
					case "antidote leaf":
						zobj.name2 = "poison antidote potion"
						zobj.color = randomgreen()
					case "resist poison bulb":
						zobj.name2 = "resist poison potion"
						zobj.color = randomgreen()
						zobj.amount = inven[invennumweeds[a]].amount
					case "resist fire root":
						zobj.name2 = "resist fire potion"
						zobj.color = randomorange()
						zobj.amount = inven[invennumweeds[a]].amount
					}

					inven[invennumbottles[a]] = xobj{}
					inven[invennumweeds[a]] = xobj{}

					if beltinvencurrentnum < len(beltinven) {
						beltinven[beltinvencurrentnum] = zobj
						findbeltinvennum()
					} else {
						inven[invennumbottles[a]] = zobj
					}
					findinvennum()
				}
			} else if len(weeds) < len(bottles) {
				for a := 0; a < len(weeds); a++ {
					zobj := xobj{}
					zobj.img = bottles[a].imgl
					zobj.name = "potion"
					zobj.color = randomcolor()
					zobj.rec = inven[invennumbottles[a]].rec
					zobj.cnt = inven[invennumbottles[a]].cnt
					zobj.kind = "potion"
					switch inven[invennumweeds[a]].name2 {
					case "health boosting pollen":
						zobj.name2 = "hp max+ potion"
						zobj.color = randomred()
					case "healing bark":
						zobj.name2 = "healing potion"
						zobj.color = randomred()
					case "cure disease seeds":
						zobj.name2 = "cure disease potion"
						zobj.color = randomyellow()
					case "antidote leaf":
						zobj.name2 = "poison antidote potion"
						zobj.color = randomgreen()
					case "resist poison bulb":
						zobj.name2 = "resist poison potion"
						zobj.color = randomgreen()
						zobj.amount = inven[invennumweeds[a]].amount
					case "resist fire root":
						zobj.name2 = "resist fire potion"
						zobj.color = randomorange()
						zobj.amount = inven[invennumweeds[a]].amount
					}

					inven[invennumbottles[a]] = xobj{}
					inven[invennumweeds[a]] = xobj{}
					if beltinvencurrentnum < len(beltinven) {
						beltinven[beltinvencurrentnum] = zobj
						findbeltinvennum()
					} else {
						inven[invennumbottles[a]] = zobj
					}
					findinvennum()
				}
			}
		} else {
			newmsg("NO INGREDIENTS >> find some weeds to make a potion")
		}
	} else {
		newmsg("NO EMPY POTION BOTTLES >> find some to make a potion")
	}

}

func makeinnerrecs() { //MARK: makeinnerrecs

	for a := 0; a < len(level); a++ {

		for b := 0; b < len(level[a].roomrec); b++ {

			level[a].roomrec[b].innerrec = level[a].roomrec[b].rec

			level[a].roomrec[b].innerrec.X += tilesize
			level[a].roomrec[b].innerrec.Y += tilesize

			level[a].roomrec[b].innerrec.Width -= tilesize * 2
			level[a].roomrec[b].innerrec.Height -= tilesize * 2

		}

	}

}
func makewater() { //MARK: makewater

	num := rInt(20, 60)

	for {

		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()

		makewaterobj(zobj, roomnumobjs)

		num--
		if num == 0 {
			break

		}
	}

}
func makewaterobj(zobj xobj, roomnumobjs int) { //MARK: makewaterobj

	zobj.name = "water"
	zobj.fixed = true
	addwater := true

	zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y-tilesize/2, tilesize, tilesize)

	for a := 0; a < len(level[roomnumobjs].objs); a++ {

		if rl.CheckCollisionRecs(level[roomnumobjs].objs[a].rec, zobj.rec) {
			addwater = false
		}
	}

	if addwater {
		x := zobj.rec.X
		y := zobj.rec.Y

		for {

			zwaterobj := rl.NewRectangle(x, y, 8, 8)

			zobj.water = append(zobj.water, zwaterobj)

			if x == zobj.rec.X {
				if flipcoin() {
					zwaterobj.X -= 8
				}
				zobj.water = append(zobj.water, zwaterobj)
			}
			if x == zobj.rec.X+(zobj.rec.Width-8) {
				if flipcoin() {
					zwaterobj.X += 8
				}
				zobj.water = append(zobj.water, zwaterobj)
			}

			if y == zobj.rec.Y {
				if flipcoin() {
					zwaterobj.Y -= 8
				}
				zobj.water = append(zobj.water, zwaterobj)
			}
			if y == zobj.rec.Y+(zobj.rec.Height-8) {
				if flipcoin() {
					zwaterobj.Y += 8
				}
				zobj.water = append(zobj.water, zwaterobj)
			}

			x += 8

			if x >= zobj.rec.X+zobj.rec.Width {
				y += 8
				x = zobj.rec.X
			}
			if y >= zobj.rec.Y+zobj.rec.Height {
				break
			}

		}

		level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)
	}

	for {

		if flipcoin() {

			addwater = true

			choose := rInt(1, 5)

			switch choose {
			case 1:
				zobj.rec.X -= tilesize
			case 2:
				zobj.rec.X += tilesize
			case 3:
				zobj.rec.Y -= tilesize
			case 4:
				zobj.rec.Y += tilesize
			}

			for a := 0; a < len(level[roomnumobjs].objs); a++ {
				if rl.CheckCollisionRecs(level[roomnumobjs].objs[a].rec, zobj.rec) {
					addwater = false
				}
			}

			checkv2 := rl.NewVector2(zobj.rec.X, zobj.rec.Y)

			if addwater {
				notinrec := true
				for a := 0; a < len(level[roomnumobjs].roomrec); a++ {

					switch choose {
					case 1, 3:
						if rl.CheckCollisionPointRec(checkv2, level[roomnumobjs].roomrec[a].rec) {
							notinrec = false
						}
					case 2:
						checkv2.X += tilesize
						if rl.CheckCollisionPointRec(checkv2, level[roomnumobjs].roomrec[a].rec) {
							notinrec = false
						}
					case 4:
						checkv2.Y += tilesize
						if rl.CheckCollisionPointRec(checkv2, level[roomnumobjs].roomrec[a].rec) {
							notinrec = false
						}
					}

				}

				if notinrec {
					addwater = false
				}
			}

			if addwater {
				x := zobj.rec.X
				y := zobj.rec.Y

				for {

					zwaterobj := rl.NewRectangle(x, y, 8, 8)

					zobj.water = append(zobj.water, zwaterobj)

					if x == zobj.rec.X {
						if flipcoin() {
							zwaterobj.X -= 8
						}
						zobj.water = append(zobj.water, zwaterobj)
					}
					if x == zobj.rec.X+(zobj.rec.Width-8) {
						if flipcoin() {
							zwaterobj.X += 8
						}
						zobj.water = append(zobj.water, zwaterobj)
					}

					if y == zobj.rec.Y {
						if flipcoin() {
							zwaterobj.Y -= 8
						}
						zobj.water = append(zobj.water, zwaterobj)
					}
					if y == zobj.rec.Y+(zobj.rec.Height-8) {
						if flipcoin() {
							zwaterobj.Y += 8
						}
						zobj.water = append(zobj.water, zwaterobj)
					}

					x += 8

					if x >= zobj.rec.X+zobj.rec.Width {
						y += 8
						x = zobj.rec.X
					}
					if y >= zobj.rec.Y+zobj.rec.Height {
						break
					}

				}

				level[roomnumobjs].objs = append(level[roomnumobjs].objs, zobj)
			}

		} else {
			break
		}
	}

}
func makebackobjs() { //MARK: makebackobjs

	num := 150

	for {

		zobj := xobj{}
		roomnumobjs := 0
		zobj.cnt, roomnumobjs = findcntr()
		zobj.color = randomgrey()

		zobj.fixed = true
		chooseimg := rInt(1, 74)
		switch chooseimg {
		case 1:
			zobj.img = weed1img
		case 2:
			zobj.img = weed2img
		case 3:
			zobj.img = weed3img
		case 4:
			zobj.img = weed4img
		case 5:
			zobj.img = weed5img
		case 6:
			zobj.img = weed6img
		case 7:
			zobj.img = weed7img
		case 8:
			zobj.img = weed8img
		case 9:
			zobj.img = weed9img
		case 10:
			zobj.img = weed10img
		case 11:
			zobj.img = back10img
		case 12:
			zobj.img = back11img
		case 13:
			zobj.img = back12img
		case 14:
			zobj.img = back13img
		case 15:
			zobj.img = back14img
		case 16:
			zobj.img = back15img
		case 17:
			zobj.img = back16img
		case 18:
			zobj.img = back17img
		case 19:
			zobj.img = back18mg
		case 20:
			zobj.img = back19img
		case 21:
			zobj.img = back1img
		case 22:
			zobj.img = back20img
		case 23:
			zobj.img = back21img
		case 24:
			zobj.img = back22img
		case 25:
			zobj.img = back23img
		case 26:
			zobj.img = back24img
		case 27:
			zobj.img = back25img
		case 28:
			zobj.img = back26img
		case 29:
			zobj.img = back27img
		case 30:
			zobj.img = back28img
		case 31:
			zobj.img = back29img
		case 32:
			zobj.img = back2img
		case 33:
			zobj.img = back30img
		case 34:
			zobj.img = back31img
		case 35:
			zobj.img = back32img
		case 36:
			zobj.img = back33img
		case 37:
			zobj.img = back34img
		case 38:
			zobj.img = back35img
		case 39:
			zobj.img = back36img
		case 40:
			zobj.img = back37img
		case 41:
			zobj.img = back38img
		case 42:
			zobj.img = back39img
		case 43:
			zobj.img = back3img
		case 44:
			zobj.img = back40img
		case 45:
			zobj.img = back41img
		case 46:
			zobj.img = back42img
		case 47:
			zobj.img = back43img
		case 48:
			zobj.img = back44img
		case 49:
			zobj.img = back45img
		case 50:
			zobj.img = back46img
		case 51:
			zobj.img = back47img
		case 52:
			zobj.img = back48img
		case 53:
			zobj.img = back49img
		case 54:
			zobj.img = back4img
		case 55:
			zobj.img = back50img
		case 56:
			zobj.img = back51img
		case 57:
			zobj.img = back52img
		case 58:
			zobj.img = back53img
		case 59:
			zobj.img = back54img
		case 60:
			zobj.img = back55img
		case 61:
			zobj.img = back56img
		case 62:
			zobj.img = back57img
		case 63:
			zobj.img = back58img
		case 64:
			zobj.img = back59img
		case 65:
			zobj.img = back5img
		case 66:
			zobj.img = back60img
		case 67:
			zobj.img = back61img
		case 68:
			zobj.img = back62img
		case 69:
			zobj.img = back63img
		case 70:
			zobj.img = back6img
		case 71:
			zobj.img = back7img
		case 72:
			zobj.img = back8img
		case 73:
			zobj.img = back9img

		}

		zobj.rec = rl.NewRectangle(zobj.cnt.X-tilesize/2, zobj.cnt.Y/2-tilesize, tilesize, tilesize)

		level[roomnumobjs].backobjs = append(level[roomnumobjs].backobjs, zobj)

		num--
		if num == 0 {
			break
		}

	}

}

func makeplayer() { //MARK: makeplayer

	player.vel = tilesize / 4
	player.velorig = player.vel
	player.intel = rInt(1, 6)
	player.luck = rInt(8, 12)
	player.str = rInt(1, 6)
	player.dex = rInt(1, 6)
	player.hp = 10
	player.hpmax = 10
	player.teleports = rInt(3, 8)
	player.rec = rl.NewRectangle(player.cnt.X-tilesize/2, player.cnt.Y-tilesize/2, tilesize, tilesize)
	player.boundrec = player.rec
	player.boundrec.X -= tilesize * 2
	player.boundrec.Y -= tilesize * 2
	player.boundrec.Width += tilesize * 4
	player.boundrec.Height += tilesize * 4
	player.roomnum = 0
	player.color = brightorange()
	player.v1 = player.cnt
	player.v1.X -= tilesize / 2
	player.v1.Y -= tilesize / 2
	player.v2 = player.v1
	player.v2.X += tilesize
	player.v3 = player.v2
	player.v3.Y += tilesize
	player.v4 = player.v3
	player.v4.X -= tilesize

}
func makemonsters() { //MARK: makemonsters

	monsters = nil
	num := 0
	monsternumlevel = currentlevelnum * 10

	for {
		zmonster := xmonster{}
		zmonster.cnt, _ = findcntr()
		zmonster.cnt.X += tilesize * 2
		zmonster.move = rInt(1, 6)
		zmonster.num = num
		zmonster.hp = rInt(1, 6)
		if flipcoin() {
			zmonster.hp += currentlevelnum
		} else {
			zmonster.hp += currentlevelnum * 2
		}

		zmonster.name = monsternames[rInt(0, len(monsternames))]

		zmonster.move = rInt(1, 5)

		zmonster.atktype = rInt(0, 5)

		choose := rInt(1, 25)

		switch choose {
		case 24: //bear
			zmonster.img = monster27img
			zmonster.frames = 4
		case 23: //snake2
			zmonster.img = monster26img
			zmonster.frames = 8
		case 22: //brown tongue
			zmonster.img = monster25img
			zmonster.frames = 4
		case 21: //pink red blob
			zmonster.img = monster24img
			zmonster.frames = 5
		case 20: //lots of eyes
			zmonster.img = monster23img
			zmonster.frames = 4
		case 19: //light blue octopus
			zmonster.img = monster22img
			zmonster.frames = 4
		case 18: //small dude
			zmonster.img = monster21img
			zmonster.frames = 4
		case 17: //purple professor
			zmonster.img = monster18img
			zmonster.frames = 4
		case 16: //blue octopus
			zmonster.img = monster17img
			zmonster.frames = 4
		case 15: //green medusa
			zmonster.img = monster16img
			zmonster.frames = 4
		case 14: //purple eyes
			zmonster.img = monster15img
			zmonster.frames = 4
		case 13: //light green blob
			zmonster.img = monster13img
			zmonster.frames = 5
		case 12: //pink eyes
			zmonster.img = monster12img
			zmonster.frames = 5
		case 11: //green antlers
			zmonster.img = monster11img
			zmonster.frames = 5
		case 10: //green eye
			zmonster.img = monster10img
			zmonster.frames = 5
		case 9: //blue horn
			zmonster.img = monster9img
			zmonster.frames = 5
		case 8: //grey blob
			zmonster.img = monster8img
			zmonster.frames = 5
		case 7: //snake
			zmonster.img = monster7img
			zmonster.frames = 5
		case 6: //eyeball
			zmonster.img = monster6img
			zmonster.frames = 5
		case 5: //brown otter
			zmonster.img = monster5img
			zmonster.frames = 5
		case 4: //blue blob
			zmonster.img = monster4img
			zmonster.frames = 4
		case 3: //pink dinosaur
			zmonster.img = monster3img
			zmonster.frames = 4
		case 2: //brown seal
			zmonster.img = monster2img
			zmonster.frames = 4
		case 1: //red
			zmonster.img = monster1img
			zmonster.frames = 4

		}
		zmonster.startx = zmonster.img.X
		zmonster.endx = zmonster.startx + (zmonster.frames * float32(16))

		switch choose {
		case 24, 25, 26:
			zmonster.endx -= 1
		}
		zmonster.rec = rl.NewRectangle(zmonster.cnt.X-tilesize/2, zmonster.cnt.Y-tilesize/2, tilesize, tilesize)
		monsters = append(monsters, zmonster)

		num++
		if num == monsternumlevel {
			break
		}
	}

}
func makeroombackg(rec rl.Rectangle) []xbackg { //MARK: makeroombackg

	var roomback []xbackg

	//floor imgs
	x := rec.X
	y := rec.Y

	for {

		zbackg := xbackg{}

		zbackg.origin = rl.NewVector2(0, 0)
		zbackg.destrec = rl.NewRectangle(x, y, tilesize, tilesize)
		zbackg.color = rl.DarkGray
		zbackg.fade = 0.2

		roomback = append(roomback, zbackg)

		x += tilesize
		if x >= rec.X+rec.Width {
			x = rec.X
			y += tilesize
		}
		if y >= rec.Y+rec.Height {
			break
		}
	}

	return roomback

}
func makeboundrec(room xroom) rl.Rectangle { //MARK: makeboundrec

	rec := rl.NewRectangle(0, 0, 0, 0)

	xleft := room.roomrec[0].rec.X
	xright := room.roomrec[0].rec.X + room.roomrec[0].rec.Width

	ytop := room.roomrec[0].rec.Y
	ybot := room.roomrec[0].rec.Y + room.roomrec[0].rec.Height

	if len(room.roomrec) > 1 {

		for a := 1; a < len(room.roomrec); a++ {

			if room.roomrec[a].rec.X < xleft {
				xleft = room.roomrec[a].rec.X
			}
			if room.roomrec[a].rec.X+room.roomrec[a].rec.Width > xright {
				xright = room.roomrec[a].rec.X + room.roomrec[a].rec.Width
			}

			if room.roomrec[a].rec.Y < ytop {
				ytop = room.roomrec[a].rec.Y
			}
			if room.roomrec[a].rec.Y+room.roomrec[a].rec.Height > ybot {
				ybot = room.roomrec[a].rec.Y + room.roomrec[a].rec.Height
			}

		}

	}

	rec = rl.NewRectangle(xleft, ytop, xright-xleft, ybot-ytop)

	return rec

}
func makecollisrec(rec rl.Rectangle) rl.Rectangle { //MARK: makecollisrec

	rec.X -= tilesize
	rec.Y -= tilesize
	rec.Width += tilesize * 2
	rec.Height += tilesize * 2

	return rec

}
func makereccnt(rec rl.Rectangle) rl.Vector2 { //MARK: makereccnt

	v2 := rl.NewVector2(0, 0)

	v2.X = rec.X
	v2.X += rec.Width / 2

	v2.Y = rec.Y
	v2.Y += rec.Height / 2

	return v2
}
func makepassage(roomnum int) { //MARK: makepassages

	choose := rInt(0, len(level[roomnum].roomrec))

	v2 := level[roomnum].roomrec[choose].cnt

	countbreak := 100

	for {

		dir := 0

		if v2.X < scrcnt.X {
			dir = 6
		} else if v2.X > scrcnt.X {
			dir = 4
		} else if v2.X == scrcnt.X {
			if v2.Y < scrcnt.Y {
				dir = 2
			} else {
				dir = 8
			}

		}
		if v2.Y < scrcnt.Y {

			if dir == 6 {
				dir = 3
			}
			if dir == 4 {
				dir = 1
			}

		} else if v2.Y > scrcnt.Y {
			if dir == 6 {
				dir = 9
			}
			if dir == 4 {
				dir = 7
			}

		}

		zroom := xroom{}
		zroom.num = len(level)
		tilemulti := float32(rInt(4, 16))
		zrec := xroomrec{}
		width := float32(3)

		switch dir {
		case 2: //down

			zrec.rec = rl.NewRectangle(v2.X, v2.Y, width*tilesize, tilemulti*tilesize)

			v2.X = zrec.rec.X - tilesize
			v2.Y = zrec.rec.Y + zrec.rec.Height - tilesize

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)

			zroom.roomrec = append(zroom.roomrec, zrec)

			zroom.boundrec = makeboundrec(zroom)

			level = append(level, zroom)
		case 8: //up

			zrec.rec = rl.NewRectangle(v2.X, v2.Y-tilemulti*tilesize, width*tilesize, tilemulti*tilesize)

			v2.X = zrec.rec.X - tilesize
			v2.Y = zrec.rec.Y

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)

			zroom.roomrec = append(zroom.roomrec, zrec)

			zroom.boundrec = makeboundrec(zroom)

			level = append(level, zroom)

		case 9: //right up
			zrec.rec = rl.NewRectangle(v2.X, v2.Y, tilemulti*tilesize, width*tilesize)
			v2.X = zrec.rec.X + zrec.rec.Width - (tilesize * 3)
			v2.Y = zrec.rec.Y + (tilesize * 3)

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)

			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec = rl.NewRectangle(v2.X, v2.Y-tilemulti*tilesize, width*tilesize, tilemulti*tilesize)

			v2.X = zrec.rec.X
			v2.Y = zrec.rec.Y

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)

			zroom.roomrec = append(zroom.roomrec, zrec)

			zroom.boundrec = makeboundrec(zroom)

			level = append(level, zroom)

		case 7: //left up
			zrec.rec = rl.NewRectangle(v2.X-tilemulti*tilesize, v2.Y, tilemulti*tilesize, width*tilesize)
			v2.X = zrec.rec.X
			v2.Y = zrec.rec.Y + (tilesize * 3)

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)

			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec = rl.NewRectangle(v2.X, v2.Y-tilemulti*tilesize, width*tilesize, tilemulti*tilesize)

			v2.X = zrec.rec.X
			v2.Y = zrec.rec.Y

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)

			zroom.roomrec = append(zroom.roomrec, zrec)

			zroom.boundrec = makeboundrec(zroom)

			level = append(level, zroom)

		case 1: //left down
			zrec.rec = rl.NewRectangle(v2.X-tilemulti*tilesize, v2.Y, tilemulti*tilesize, width*tilesize)
			v2.X = zrec.rec.X
			v2.Y = zrec.rec.Y + (tilesize * 3)

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)

			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec = rl.NewRectangle(v2.X, v2.Y, width*tilesize, tilemulti*tilesize)

			v2.X = zrec.rec.X + tilesize*3
			v2.Y = zrec.rec.Y + zrec.rec.Height - (tilesize * 3)

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)

			zroom.roomrec = append(zroom.roomrec, zrec)

			zroom.boundrec = makeboundrec(zroom)

			level = append(level, zroom)

		case 3: //right down
			zrec.rec = rl.NewRectangle(v2.X, v2.Y, tilemulti*tilesize, width*tilesize)
			v2.X = zrec.rec.X + zrec.rec.Width - (tilesize * 3)
			v2.Y = zrec.rec.Y

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)

			zroom.roomrec = append(zroom.roomrec, zrec)

			zrec.rec = rl.NewRectangle(v2.X, v2.Y, width*tilesize, tilemulti*tilesize)

			v2.X = zrec.rec.X
			v2.Y = zrec.rec.Y + zrec.rec.Height - (tilesize * 3)

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)

			zroom.roomrec = append(zroom.roomrec, zrec)

			zroom.boundrec = makeboundrec(zroom)

			level = append(level, zroom)

		case 4: //left
			zrec.rec = rl.NewRectangle(v2.X-tilemulti*tilesize, v2.Y, tilemulti*tilesize, width*tilesize)
			v2.X = zrec.rec.X + tilesize
			v2.Y = zrec.rec.Y + tilesize

			zrec.cnt = makereccnt(zrec.rec)
			zrec.collisrec = makecollisrec(zrec.rec)

			zroom.roomrec = append(zroom.roomrec, zrec)
			level = append(level, zroom)

		case 6: //right
			zrec.rec = rl.NewRectangle(v2.X, v2.Y, tilemulti*tilesize, width*tilesize)
			v2.X = zrec.rec.X + zrec.rec.Width - tilesize
			v2.Y = zrec.rec.Y + tilesize

			zroom.roomrec = append(zroom.roomrec, zrec)

			zroom.boundrec = makeboundrec(zroom)

			level = append(level, zroom)
		}

		end := false
		for a := 0; a < origlevellen; a++ {
			if a != roomnum {
				for b := 0; b < len(level[a].roomrec); b++ {
					if rl.CheckCollisionRecs(zroom.roomrec[0].rec, level[a].roomrec[b].rec) {
						end = true
					}
				}
			}
		}

		if end {
			break
		}

		countbreak--
		if countbreak == 0 {
			break
		}

		if dir == 0 {
			break
		}
	}

}
func makefx() { //MARK: makefx

	y := float32(0)

	for {
		zscan := xscanline{}
		zscan.v1 = rl.NewVector2(0, y)
		zscan.v2 = rl.NewVector2(scrwf32, y)

		scanlines = append(scanlines, zscan)
		y += 7

		if y > scrhf32 {
			break
		}
	}

}

func makeinven() { //MARK: makeinven

	num := 28

	for a := 0; a < num; a++ {

		zinven := xobj{}
		zinven.rec = rl.NewRectangle(0, 0, tilesize, tilesize)
		zinven.color = randomcolor()
		inven = append(inven, zinven)

	}

	num = 20

	for a := 0; a < num; a++ {

		zinven := xobj{}
		zinven.rec = rl.NewRectangle(0, 0, tilesize, tilesize)
		zinven.color = randomcolor()
		beltinven = append(beltinven, zinven)

	}
}

func makenewobj(objtype, roomnum int, cnt rl.Vector2) { //MARK: makenewobj

	zobj := xobj{}

	switch objtype {
	case 3: //coin

		zobj.img = coinimg
		zobj.name = "coin"
		zobj.collect = true
		zobj.color = rl.White
		zobj.fixed = true

	case 2: //grave obj
		if flipcoin() {
			choose := rInt(0, len(jewelryimgs))
			zobj.img = jewelryimgs[choose]
			zobj.color = brightyellow()
			zobj.collect = true
			zobj.name = "jewelry"
			zobj.amount = rInt(5, 21)
			zobj.usetype = rInt(1, 10)
			zobj.kind = "jewel"
			switch zobj.usetype {
			case 1:
				zobj.name2 = "jewelry of dexterity +" + fmt.Sprint(zobj.amount)
			case 2:
				zobj.name2 = "jewelry of strength +" + fmt.Sprint(zobj.amount)
			case 3:
				zobj.name2 = "jewelry of luck +" + fmt.Sprint(zobj.amount)
			case 4:
				zobj.name2 = "jewelry of intelligence +" + fmt.Sprint(zobj.amount)
			case 5:
				zobj.name2 = "jewelry of fire protection +" + fmt.Sprint(zobj.amount*10) + "%"
			case 6:
				zobj.name2 = "jewelry of poison protection +" + fmt.Sprint(zobj.amount*10) + "%"
			case 7:
				zobj.name2 = "jewelry of disease immunity"
			case 8:
				zobj.name2 = "jewelry of max hp +" + fmt.Sprint(zobj.amount)
			case 9:
				zobj.name2 = "jewelry of teleport"
				zobj.timer = rInt32(fpsint*5, fpsint*10)
			}
		} else {
			choose := rInt(0, len(gemimgs))
			zobj.img = gemimgs[choose]
			zobj.color = randomcolor()
			zobj.collect = true
			zobj.amount = rInt(10, 21)
			zobj.kind = "gem"
			zobj.name = "gem"
			zobj.name2 = "gem value is " + fmt.Sprint(zobj.amount)

		}
	case 1: //chest obj

		newobj := makeopenchestobj()

		zobj = newobj
		zobj.color = randomcolor()
	}

	breakcount := 100
	for {

		zobj.rec = rl.NewRectangle(cnt.X+rFloat32(-tilesize*4, tilesize*4), cnt.Y+rFloat32(-tilesize*4, tilesize*4), tilesize, tilesize)

		collides := false

		for a := 0; a < len(level[roomnum].objs); a++ {
			if rl.CheckCollisionRecs(zobj.rec, level[roomnum].objs[a].rec) {
				collides = true
			}
		}

		if rl.CheckCollisionRecs(zobj.rec, player.rec) {
			collides = true
		}

		collidesborders := false

		if !collides {

			zobj.v1 = rl.NewVector2(zobj.rec.X, zobj.rec.Y)
			zobj.v2 = rl.NewVector2(zobj.rec.X+tilesize, zobj.rec.Y)
			zobj.v3 = rl.NewVector2(zobj.rec.X+tilesize, zobj.rec.Y+tilesize)
			zobj.v4 = rl.NewVector2(zobj.rec.X, zobj.rec.Y+tilesize)

			check1, check2, check3, check4 := false, false, false, false

			for a := 0; a < len(level[roomnum].roomrec); a++ {
				if rl.CheckCollisionPointRec(zobj.v1, level[roomnum].roomrec[a].rec) {
					check1 = true
				}
				if rl.CheckCollisionPointRec(zobj.v2, level[roomnum].roomrec[a].rec) {
					check2 = true
				}
				if rl.CheckCollisionPointRec(zobj.v3, level[roomnum].roomrec[a].rec) {
					check3 = true
				}
				if rl.CheckCollisionPointRec(zobj.v4, level[roomnum].roomrec[a].rec) {
					check4 = true
				}
			}

			if check1 && check2 && check3 && check4 {
				collidesborders = true
			}
		}

		if collidesborders {
			break
		}

		breakcount--
		if breakcount == 0 {
			break
		}
	}

	zobj.cnt = rl.NewVector2(zobj.rec.X+tilesize/2, zobj.rec.Y+tilesize/2)

	zobj.boundrec = zobj.rec
	zobj.boundrec.X -= tilesize * 2
	zobj.boundrec.Y -= tilesize * 2
	zobj.boundrec.Width += tilesize * 4
	zobj.boundrec.Height += tilesize * 4

	level[roomnum].objs = append(level[roomnum].objs, zobj)

}
func makeopenchestobj() xobj { //MARK: makeopenchestobj

	zobj := xobj{}

	choose := rInt(1, 14)

	switch choose {
	case 13: //empty potion
		zobj.collect = true
		choose := rInt(0, len(potionemptyimgs))
		zobj.img = potionemptyimgs[choose]
		zobj.imgl = potionimgs[choose]
		zobj.color = randombluelight()
		zobj.name = "empty potion bottle"
	case 12: //bow
		choose := rInt(0, len(bowimgs))
		zobj.img = bowimgs[choose]
		zobj.origimg = zobj.img
		zobj.imgl = bowimgsl[choose]
		zobj.name = "bow"
		zobj.kind = "weap"
		zobj.collect = true
		zobj.vel1 = 5
		zobj.atk = 1 + currentlevelnum
		if rolldice() > 4 {
			zobj.name2 = "bow of speed +" + fmt.Sprint(zobj.atk) + " damage"
			zobj.vel1 = rFloat32(6, 10)
		}
	case 11: //crossbow
		choose := rInt(0, len(crossbowimgs))
		zobj.img = crossbowimgs[choose]
		zobj.origimg = zobj.img
		zobj.imgl = crossbowimgsl[choose]
		zobj.name = "crossbow"
		zobj.kind = "weap"
		zobj.collect = true
		zobj.vel1 = 10
		zobj.atk = 2 + currentlevelnum
		if rolldice() > 4 {
			zobj.name2 = "crossbow of speed"
			zobj.vel1 = rFloat32(12, 18)
		}
		if rolldice() > 4 {
			if zobj.name2 == "crossbow of speed" {
				zobj.name2 = "heavy crossbow of speed +" + fmt.Sprint(zobj.atk) + " damage"
			} else {
				zobj.name2 = "heavy crossbow +" + fmt.Sprint(zobj.atk) + " damage"
			}
			zobj.atk += rInt(2, 5)
		}
	case 10: //spear
		choose := rInt(0, len(spearimgs))
		zobj.img = spearimgs[choose]
		zobj.origimg = zobj.img
		zobj.imgl = spearimgsl[choose]
		zobj.name = "spear"
		zobj.kind = "weap"
		zobj.vel1 = 8
		zobj.collect = true
		zobj.atk = 2 + currentlevelnum
		zobj.numberof = rInt(50, 100)
		if rolldice() > 4 {
			zobj.name2 = "spear of speed +" + fmt.Sprint(zobj.atk) + " damage"
			zobj.vel1 = rFloat32(10, 15)
		}
	case 9: //throwing axe
		choose := rInt(0, len(throwingaxeimgs))
		zobj.img = throwingaxeimgs[choose]
		zobj.origimg = zobj.img
		zobj.imgl = throwingaxeimgsl[choose]
		zobj.name = "throwing axe"
		zobj.kind = "weap"
		zobj.vel1 = 8
		zobj.collect = true
		zobj.rotates = true
		zobj.atk = 1 + currentlevelnum
		zobj.numberof = rInt(50, 100)
		if rolldice() > 4 {
			zobj.name2 = "throwing axe of speed +" + fmt.Sprint(zobj.atk) + " damage"
			zobj.vel1 = rFloat32(10, 15)
		}
	case 8: //axe
		choose := rInt(0, len(axeimgs))
		zobj.img = axeimgs[choose]
		zobj.origimg = zobj.img
		zobj.imgl = axeimgsl[choose]
		zobj.name = "axe"
		zobj.kind = "weap"
		zobj.collect = true
		zobj.atk = 3 + currentlevelnum
		zobj.meleerange = tilesize * 4
		zobj.meleerangerec = makerangerec(zobj)
	case 7: //wand
		choose := rInt(0, len(wandimgs))
		zobj.img = wandimgs[choose]
		zobj.origimg = zobj.img
		zobj.imgl = wandimgsl[choose]
		zobj.name = "wand"
		zobj.kind = "weap"
		zobj.collect = true
		zobj.atk = 2 + currentlevelnum
		zobj.numberof = rInt(50, 100)

		zobj.ability = rInt(1, 5)

		switch zobj.ability {
		case 4:
			zobj.name2 = "wand of chaining light thingy"
		case 3:
			zobj.name2 = "wand of froggy rain"
		case 2:
			zobj.name2 = "wand of goes round thingy"
		case 1:
			zobj.name2 = "wand of green laser thingy"
		}
	case 6: //mace
		choose := rInt(0, len(maceimgs))
		zobj.img = maceimgs[choose]
		zobj.origimg = zobj.img
		zobj.imgl = maceimgsl[choose]
		zobj.name = "mace"
		zobj.kind = "weap"
		zobj.collect = true
		zobj.atk = 2 + currentlevelnum
		zobj.meleerange = tilesize * 3
		zobj.meleerangerec = makerangerec(zobj)
	case 5: //scythe
		choose := rInt(0, len(scytheimgs))
		zobj.img = scytheimgs[choose]
		zobj.origimg = zobj.img
		zobj.imgl = scytheimgsl[choose]
		zobj.name = "scythe"
		zobj.kind = "weap"
		zobj.collect = true
		zobj.atk = 2 + currentlevelnum
		zobj.meleerange = tilesize * 5
		zobj.meleerangerec = makerangerec(zobj)
	case 4: //club
		choose := rInt(0, len(clubimgs))
		zobj.img = clubimgs[choose]
		zobj.origimg = zobj.img
		zobj.imgl = clubimgsl[choose]
		zobj.name = "club"
		zobj.kind = "weap"
		zobj.collect = true
		zobj.atk = 2 + currentlevelnum
		zobj.meleerange = tilesize * 3
		zobj.meleerangerec = makerangerec(zobj)
	case 3: //dagger
		choose := rInt(0, len(daggerimgs))
		zobj.img = daggerimgs[choose]
		zobj.origimg = zobj.img
		zobj.imgl = daggerimgsl[choose]
		zobj.name = "dagger"
		zobj.kind = "weap"
		zobj.collect = true
		zobj.atk = 1 + currentlevelnum
		zobj.meleerange = tilesize * 2
		zobj.meleerangerec = makerangerec(zobj)
	case 2: //sword
		choose := rInt(0, len(swordimgs))
		zobj.img = swordimgs[choose]
		zobj.origimg = zobj.img
		zobj.imgl = swordimgsl[choose]
		zobj.name = "sword"
		zobj.kind = "weap"
		zobj.collect = true
		zobj.atk = 2 + currentlevelnum
		zobj.meleerange = tilesize * 3
		zobj.meleerangerec = makerangerec(zobj)
	case 1: //ninja star
		zobj.img = ninjastar1img
		zobj.name = "ninja star"
		zobj.kind = "weap"
		zobj.vel1 = 12
		zobj.collect = true
		zobj.rotates = true
		zobj.atk = 1 + currentlevelnum
		zobj.numberof = rInt(50, 100)
	}

	return zobj

}
func makededmonster(v2 rl.Vector2) { //MARK: makededmonster

	zded := xdedmonster{}
	zded.v2 = v2
	zded.timer = fps

	num := rInt(10, 20)

	for a := 0; a < num; a++ {

		zcirc := xcircle{}
		zcirc.color = randomred()
		zcirc.rad = rFloat32(30, 60)
		zcirc.v2 = zded.v2
		zcirc.v2.X += rFloat32(-tilesize*2, tilesize*2)
		zcirc.v2.Y += rFloat32(-tilesize*2, tilesize*2)

		zded.circles = append(zded.circles, zcirc)
	}

	dedmonsters = append(dedmonsters, zded)

	killcount++

	if killcount == monsternumlevel {
		clearedlevellootnum = rInt(3, 8)
		clearedleveltimer = fps * 4
		clearedmonsters = true
	}

	monsterkills++
	if monsterkills == nextbossnum {
		makeboss()
		nextbossnum += rInt(5, 10)
	}

	score += currentlevelnum * 10

}

// MARK: CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE
func initial() { //MARK: initial

	musicon = true
	soundfxon = true

	rl.SetMasterVolume(soundvol)

	diedy = -scrhf32
	deathon = true

	introplayx = scrwf32 - tilesize
	introtxtx, introtxtx2 = -200, -700
	introtxty = -scrh / 2

	pigeononoff = true
	pigeontimer = fps * rInt32(20, 100)

	teleporton = true
	borderrec = rl.NewRectangle(scrwf32/2-defw/2, scrhf32/2-defh/2, defw, defh)
	camera.Zoom = 1.5
	cammap.Zoom = 0.1
	invenrec = rl.NewRectangle(0, 0, tilesize*3, scrhf32)
	statsrec = rl.NewRectangle(scrwf32-tilesize*3, 0, tilesize*3, scrhf32)
	msgrec = rl.NewRectangle(invenrec.X+invenrec.Width, 0, scrwf32-statsrec.Width-invenrec.Width, tilesize+tilesize/8)
	footerrec = msgrec
	footerrec.Y = scrhf32 - footerrec.Height
	selpoint = blankv2
	playeranimatetimer = rInt(10, 21)
	playeremotetimer = rInt(120, 360)
	scan = true
	ghost = true
	activammonum = blankint
	activweaponnum = blankint
	playeremoteon = false

	makeimgs()
	makesounds()
	makelevel()
	makeplayer()
	makepet()
	makefx()
	makeinven()
	makeshop()

	newmsg("the higher the dungeon level the more difficult it becomes so choose wisely before descending")

	introon = true
	pause = true

}
func timers() { //MARK: timers

	if mouseclicknum > 0 && !rl.IsMouseButtonDown(rl.MouseRightButton) {
		mouseclicknum = 0
	}

	if weaponrangetimer > 0 {
		weaponrangetimer--
	} else {
		weaponrangeon = false
	}

	if pigeononoff {
		if pigeontimer > 0 && !pigeonon {
			pigeontimer--
			if pigeontimer == 1 {
				chosepigeonphrase = false
			}
		} else {

			if !chosepigeonphrase {
				pigeonphrase = pigeonwords[rInt(0, len(pigeonwords))]
				chosepigeonphrase = true
			}
			pigeonon = true

		}
	}

	if frograintimer > 0 {
		frograintimer--
	}

	if destroybelttimer > 0 {
		destroybelttimer--
	} else {
		destroybeltitemon = false
		destroybeltnum = blankint
	}
	if destroytimer > 0 {
		destroytimer--
	} else {
		destroyitemon = false
		destroynum = blankint
	}

	if clearedleveltimer > 0 {
		clearedleveltimer--
	} else {
		clearedmonsters = false
	}

	if player.damppause > 0 {
		player.damppause--
	}

	if player.damptimer > 0 {
		player.damptimer--
		if player.damptimer == 0 {
			player.dampcount = 0
		}
	}

	if player.poisoned {

		player.poisontimer--

		if player.poisontimer == 0 {
			player.poisoned = false
		}

		if player.poisontimer%fps == 0 {
			player.hp--
			if player.hp < 0 {
				player.hp = 0
			}
			if soundfxon && !mute {
				rl.PlaySoundMulti(playerdamageaud)
			}
		}

	}
	if player.burning {

		player.burntimer--

		if player.burntimer == 0 {
			player.burning = false
		}

		if player.burntimer%fps == 0 {
			player.hp--
			if player.hp < 0 {
				player.hp = 0
			}
			if soundfxon && !mute {
				rl.PlaySoundMulti(playerdamageaud)
			}
		}

	}
	if player.sick {

		player.sicktimer--

		if player.sicktimer == 0 {
			player.sick = false
		}

		if player.sicktimer%fps == 0 {
			player.hp--
			if player.hp < 0 {
				player.hp = 0
			}
			if soundfxon && !mute {
				rl.PlaySoundMulti(playerdamageaud)
			}
		}

	}
	if player.rainbow {
		player.color = randomcolor()
	} else if player.poisoned && player.sick && player.burning {

		choose := rolldice()
		if choose < 3 {
			player.color = randomgreen()
		} else if choose > 2 && choose < 5 {
			player.color = randomyellow()
		} else if choose > 4 {
			player.color = randomorange()
		}
	} else if player.poisoned && player.sick {
		if flipcoin() {
			player.color = randomgreen()
		} else {
			player.color = randomyellow()
		}
	} else if player.poisoned && player.burning {
		if flipcoin() {
			player.color = randomgreen()
		} else {
			player.color = randomorange()
		}
	} else if player.sick && player.burning {
		if flipcoin() {
			player.color = randomyellow()
		} else {
			player.color = randomorange()
		}
	} else if player.poisoned {
		player.color = randomgreen()
	} else if player.sick {
		player.color = randomyellow()
	} else if player.burning {
		player.color = randomorange()
	} else {
		player.color = brightorange()
	}

	if playeremotetimer > 0 {
		playeremotetimer--
	} else {
		if playeremoteon {
			playeremoteon = false
		} else {
			playeremoteon = true
			emoteimgnum = rInt(0, len(emoteimgs))
			emoteimg = emoteimgs[emoteimgnum]
			emoteimgx = emoteimg.X
			emoteselected = true
		}
		playeremotetimer = rInt(120, 360)
	}

	if newmsgtimer > 0 {
		newmsgtimer--
	}

	if magictimer > 0 {
		magictimer--
		if magictimer == 0 {
			magicon = false
			activmagic = nil
		}
	}

	if weaponro1 {
		weaponrotimer--
		if weaponrotimer == 0 {
			weaponro1 = false
			weaponro2 = true
			weaponrotimer = weaponrotime
		}

	} else if weaponro2 {
		weaponrotimer--
		if weaponrotimer == 0 {
			weaponro2 = false
			weaponro3 = true
			weaponrotimer = weaponrotime
		}
	} else if weaponro3 {
		weaponrotimer--
		if weaponrotimer == 0 {
			weaponro3 = false
			weaponrotimer = 0
		}
	}

	if player.hpppause > 0 {
		player.hpppause--
	}

	if frames%1 == 0 {
		if onoff1 {
			onoff1 = false
		} else {
			onoff1 = true
		}
	}

	if frames%2 == 0 {
		if onoff2 {
			onoff2 = false
		} else {
			onoff2 = true
		}
	}
	if frames%3 == 0 {
		if onoff3 {
			onoff3 = false
		} else {
			onoff3 = true
		}
	}
	if frames%6 == 0 {
		if onoff6 {
			onoff6 = false
		} else {
			onoff6 = true
		}
	}
	if frames%10 == 0 {
		if onoff10 {
			onoff10 = false
		} else {
			onoff10 = true
		}
	}
	if frames%15 == 0 {
		if onoff15 {
			onoff15 = false
		} else {
			onoff15 = true
		}
	}
	if frames%30 == 0 {
		if onoff30 {
			onoff30 = false
		} else {
			onoff30 = true
		}
	}
	if frames%60 == 0 {
		if onoff60 {
			onoff60 = false
		} else {
			onoff60 = true
		}
	}

}
func timers2() { //MARK: timers2

	if clickpause > 0 {
		clickpause--
	}

	if introtimer > 0 {
		introtimer--
	}

	if !pause {
		runtimer++
		if runtimer%60 == 0 {
			upruntime()
		}
	}

	if nokeystimer > 0 {
		nokeystimer--
		if nokeystimer == 1 {
			nokeys = false
		}
	}

	if frames%4 == 0 {
		coinimg.X += 16
		if coinimg.X > 65 {
			coinimg.X = 1
		}
	}
	if fadeblinkon {
		if fadeblink > 0.4 {
			fadeblink -= 0.05
		} else {
			fadeblinkon = false
		}
	} else {
		if fadeblink < 0.9 {
			fadeblink += 0.05
		} else {
			fadeblinkon = true
		}
	}
}
func closewinloc(x, y float32, col, col2 rl.Color) bool { //MARK: closewinloc

	close := false

	rec := rl.NewRectangle(x, y, tilesize, tilesize)
	if rl.CheckCollisionPointRec(mousev2, rec) {
		rl.DrawRectangleRec(rec, col2)
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			close = true
		}
	} else {
		rl.DrawRectangleRec(rec, col)
	}
	return close

}
func inp() { //MARK: inp

	if rl.IsKeyPressed(rl.KeyF3) {
		clearedlevellootnum = rInt(4, 8)
		killcount = monsternumlevel
		clearedleveltimer = fps * 4
		clearedmonsters = true
	}

	if rl.IsKeyPressed(rl.KeyH) {
		if helpon {
			helpon = false
			pause = false
		} else {
			helpon = true
			pause = true
		}
	}

	if rl.IsKeyPressed(rl.KeyL) {
		if displaymsgs {
			displaymsgs = false
			pause = false
		} else {
			displaymsgs = true
			pause = true
		}
	}
	if rl.IsKeyPressed(rl.KeyM) {
		if mapon {
			mapon = false
			pause = false
		} else {
			mapon = true
			pause = true
			cammap.Target = player.cnt
			cammap.Offset.X = scrwf32 / 2
			cammap.Offset.Y = scrhf32 / 2
		}
	}

	if rl.IsKeyPressed(rl.KeyPause) {
		if pause {
			pause = false
		} else {
			pause = true
		}
	}

	if rl.IsKeyPressed(rl.KeyF1) {
		if dev {
			dev = false
		} else {
			dev = true
		}
	}

}
func scr(num int) { //MARK: scr

	switch num {
	case 0:
		rl.InitWindow(0, 0, "")
		scrhint = rl.GetScreenHeight()
		scrwint = rl.GetScreenWidth()
		rl.CloseWindow()
	}

	scrhf32 = float32(scrhint)
	scrwf32 = float32(scrwint)
	scrh = int32(scrhint)
	scrw = int32(scrwint)

	scrcnt = rl.NewVector2(scrwf32/2, scrhf32/2)

	camera.Zoom = 1.0

	if scrw < 1920 {
		multi = float32(2)
		tilesize = basetile * multi
		txtdef = txts
		txtl = 20
		txtl2 = 30
	}

}
func raylib() { //MARK: raylib
	rl.SetConfigFlags(rl.FlagMsaa4xHint) // enable 4X anti-aliasing

	rl.InitWindow(scrw, scrh, "the endless dungeons of pixel")

	rl.InitAudioDevice()

	//rl.SetWindowSize(scrwint, scrhint)
	//rl.SetWindowPosition(0, 681)

	//rl.ToggleFullscreen()

	rl.HideCursor()
	imgs = rl.LoadTexture("data/imgs.png") // load images

	initial()
	rl.SetTargetFPS(fps)

	for !rl.WindowShouldClose() {
		frames++

		if musicon && !introon && !mute {
			rl.PlayMusicStream(levmusic)
			rl.UpdateMusicStream(levmusic)
		}

		mousev2 = rl.GetMousePosition()
		mousev2world = rl.GetScreenToWorld2D(mousev2, camera)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		nocambackg()

		rl.BeginMode2D(camera)
		cam()

		rl.EndMode2D()
		if pigeononoff {
			if pigeonon {
				drawpigeon()
			}
		}
		nocam()

		if mapon {
			drawmap()
		}
		if shopon {
			drawshop()
		}
		if died && !scoreon {
			drawdied()
		}
		if scoreon {
			drawscore()
		}
		if settingson {
			drawsettings()
		}
		if introon {
			drawintro()
		}
		if helpon {
			drawhelp()
		}

		if endgamewindow {
			drawendgame()
			rl.SetExitKey(rl.KeyY)
		} else {
			rl.SetExitKey(rl.KeyEscape)
		}

		nocamMap()
		if dev {
			devui()
		}
		//cursor
		if !introon {
			cursorv1 := rl.NewVector2(mousev2.X+20, mousev2.Y+5)
			cursorv2 := rl.NewVector2(mousev2.X+5, mousev2.Y+20)
			shadow1 := cursorv1
			shadow1.Y += 5
			shadow1.X -= 2
			shadow2 := cursorv2
			shadow2.Y += 5
			shadow2.X -= 2
			shadow3 := mousev2
			shadow3.Y += 5
			shadow3.X -= 2

			rl.DrawTriangle(shadow1, shadow3, shadow2, rl.Fade(rl.Black, 0.5))
			rl.DrawTriangle(cursorv1, mousev2, cursorv2, rl.Fade(rl.Magenta, fadeblink))
			rl.DrawTriangleLines(cursorv1, mousev2, cursorv2, rl.Black)
		}
		rl.EndDrawing()
		update()
	}

	rl.StopSoundMulti()
	rl.CloseAudioDevice()

	rl.CloseWindow()

}
func main() { //MARK: main
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLog(rl.LogError)      // hides info window
	scr(0)
	raylib()
}
func txtarrow(txt string, rec rl.Rectangle) { //MARK: txtarrow

	txtlen := rl.MeasureText(txt, txtdef)
	cntr := rl.NewVector2(rec.X+rec.Width/2, rec.Y-tilesize)

	if onoff10 {
		cntr.Y += 10
	}

	backrec := rl.NewRectangle((cntr.X-float32(txtlen/2))-10, cntr.Y-2, float32(txtlen+20), float32(txtdef)+4)
	backrec2 := rl.NewRectangle((cntr.X-float32(txtlen/2))-12, cntr.Y, float32(txtlen+20), float32(txtdef)+4)

	tri1 := rl.NewVector2(cntr.X-10, cntr.Y)
	tri2 := rl.NewVector2(cntr.X+10, cntr.Y)
	tri3 := rl.NewVector2(cntr.X, cntr.Y+tilesize)

	rl.DrawRectangleRec(backrec2, randomcolor())
	rl.DrawRectangleRec(backrec, rl.White)

	rl.DrawTriangle(tri3, tri2, tri1, rl.White)

	rl.DrawText(txt, int32(cntr.X-float32(txtlen/2))-1, int32(cntr.Y), txtdef, randomcolor())
	rl.DrawText(txt, int32(cntr.X-float32(txtlen/2)), int32(cntr.Y-1), txtdef, rl.Black)
}
func txthere(txt string, x, y float32) { //MARK: txthere

	txtlen := rl.MeasureText(txt, txtdef)

	rec := rl.NewRectangle(x, y, float32(txtlen+int32(tilesize)), float32(txtdef+(txtdef/2)))

	rec2 := rec
	rec2.X -= 4
	rec2.Y += 4

	rl.DrawRectangleRec(rec2, rl.Black)
	rl.DrawRectangleRec(rec, brightorange())

	rl.DrawText(txt, rec.ToInt32().X+rec.ToInt32().Width/2-txtlen/2, rec.ToInt32().Y+txtdef/4, txtdef, rl.Black)

}

// MARK: colors
// https://www.rapidtables.com/web/color/RGB_Color.html
func darkred() rl.Color {
	color := rl.NewColor(55, 0, 0, 255)
	return color
}
func semidarkred() rl.Color {
	color := rl.NewColor(70, 0, 0, 255)
	return color
}
func brightorange() rl.Color {
	color := rl.NewColor(253, 95, 0, 255)
	return color
}
func brightred() rl.Color {
	color := rl.NewColor(230, 0, 0, 255)
	return color
}
func randomgrey() rl.Color {
	color := rl.NewColor(uint8(rInt(160, 193)), uint8(rInt(160, 193)), uint8(rInt(160, 193)), 255)
	return color
}
func randombluelight() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 180)), uint8(rInt(120, 256)), uint8(rInt(120, 256)), 255)
	return color
}
func randombluedark() rl.Color {
	color := rl.NewColor(0, 0, uint8(rInt(120, 250)), 255)
	return color
}
func randomyellow() rl.Color {
	color := rl.NewColor(255, uint8(rInt(150, 256)), 0, 255)
	return color
}
func randomorange() rl.Color {
	color := rl.NewColor(uint8(rInt(250, 256)), uint8(rInt(60, 210)), 0, 255)
	return color
}
func randomred() rl.Color {
	color := rl.NewColor(uint8(rInt(128, 256)), uint8(rInt(0, 129)), uint8(rInt(0, 129)), 255)
	return color
}
func randomgreen() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 170)), uint8(rInt(100, 256)), uint8(rInt(0, 50)), 255)
	return color
}
func randomcolor() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
	return color
}
func brightyellow() rl.Color {
	color := rl.NewColor(uint8(255), uint8(255), uint8(0), 255)
	return color
}
func brightbrown() rl.Color {
	color := rl.NewColor(uint8(218), uint8(165), uint8(32), 255)
	return color
}
func brightgrey() rl.Color {
	color := rl.NewColor(uint8(212), uint8(212), uint8(213), 255)
	return color
}

// MARK: random numbers
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	i := rand.Intn(max-min) + min
	return int32(i)
}
func rFloat32(min, max float32) float32 {
	return (rand.Float32() * (max - min)) + min
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}

// MARK: other functions
func orbitpoint(cnt, point rl.Vector2, angle float32) rl.Vector2 {

	//THIS DOES NOT WORK AS A FUNCTION COPY & PASTE TO RELEVANT PART

	angle = angle * (math.Pi / 180)

	newx := float32(math.Cos(float64(angle)))*(point.X-cnt.X) - float32(math.Sin(float64(angle)))*(point.Y-cnt.Y) + cnt.X

	newy := float32(math.Sin(float64(angle)))*(point.X-cnt.X) + float32(math.Cos(float64(angle)))*(point.Y-cnt.Y) + cnt.Y

	point2 := rl.NewVector2(newx, newy)

	return point2

}
func lengthtwopoints(v1, v2 rl.Vector2) float32 {

	num := float32(0)

	x2 := getabs(v2.X)
	x1 := getabs(v1.X)

	y2 := getabs(v2.Y)
	y1 := getabs(v1.Y)

	num = squareroot(squarenum(x2-x1) + squarenum(y2-y1))

	return num
}
func squareroot(num float32) float32 {
	num = float32(math.Sqrt(float64(num)))
	return num
}
func squarenum(num float32) float32 { // num*num
	num = num * num
	return num
}
func lastdigits(num int) int {
	number := num % 1e2 //change 1e2 to 1e3 to 1e4 etc for more digit places
	return number
}
func firstdigits(num int) int {
	number := num / 1e3 //change 1e2 to 1e3 to 1e4 etc for more digit places
	return number
}
func timehere(x, y float32) {
	currentTime := time.Now()
	txtlen := rl.MeasureText(currentTime.Format("15:04"), txtdef)
	x -= float32(txtlen + txtdef)
	rl.DrawText(currentTime.Format("15:04"), int32(x), int32(y), txtdef, rl.White)
}
func getabs(num float32) float32 {
	return float32(math.Abs(float64(num)))
}
func absdiff32(num, num2 float32) float32 {

	diff := float32(0)
	if num == num2 {
		diff = 0
	} else if num == 0 || num2 == 0 {
		if num == 0 {
			diff = float32(math.Abs(float64(num2)))
		} else {
			diff = float32(math.Abs(float64(num)))
		}
	} else if num > 0 && num2 > 0 {
		if num > num2 {
			diff = num - num2
		} else {
			diff = num2 - num
		}
	} else if num > 0 && num2 < 0 || num < 0 && num2 > 0 {

		if num > 0 {
			diff = num + float32(math.Abs(float64(num2)))
		} else {
			diff = num2 + float32(math.Abs(float64(num)))
		}

	} else if num < 0 && num2 < 0 {
		num = float32(math.Abs(float64(num)))
		num2 = float32(math.Abs(float64(num2)))
		if num > num2 {
			diff = num - num2
		} else {
			diff = num2 - num
		}
	}

	return diff

}
func angle2points(start, destination rl.Vector2) float32 { //make sure destination vector is vec2
	angle := float32(math.Atan2(float64(destination.Y-start.Y), float64(destination.X-start.X)))*(180/math.Pi) + 90
	//change +30 (addition value at end) to angle to compensate for polygon rotation difference
	return angle

}

func remobj(s []xobj, index int) []xobj { //remove struct from a slice
	return append(s[:index], s[index+1:]...)
}
func remdedmons(s []xdedmonster, index int) []xdedmonster { //remove struct from a slice
	return append(s[:index], s[index+1:]...)
}
func remcirc(s []xcircle, index int) []xcircle { //remove struct from a slice
	return append(s[:index], s[index+1:]...)
}

func remstring(s []string, index int) []string { //remove string from a slice
	return append(s[:index], s[index+1:]...)
}
func diagsquare(sidelength float32) float32 {
	return sidelength * float32(math.Sqrt(2))
}
func circlearea(radius float32) float32 {
	return math.Pi * radius * radius
}
func gcd(num1, num2 float32) float32 {

	num164 := float64(num1)
	num264 := float64(num2)

	//Calculate GCD
	num3 := math.Mod(num164, num264)

	for num3 > 0 {

		num164 = num264
		num264 = num3
		num3 = math.Mod(num164, num264)

	}

	return float32(num264)

}
