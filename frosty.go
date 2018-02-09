package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ttacon/emoji"
)

// frostyHandler responds to frosty command with some random stuff
func frostyHandler(s *discordgo.Session, m *discordgo.MessageCreate, command commandWithArgs) {
	var err error
	if len(command.args) == 2 && command.args[0] == "or" && command.args[1] == "else" {
		// Special for TngB (actually - that whole command is for him :rolling_eye:)
		_, err = s.ChannelMessageSend(m.ChannelID, ":snowman2: :gun:")
	} else {
		_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(":snowman2: %s", getRandomEmoji()))
	}

	if err != nil {
		fmt.Println(err)
	}
}

// getRandomEmoji gets the :X: notation for a random emoji
func getRandomEmoji() string {
	if allEmojis == nil || len(allEmojis) == 0 {
		return ":gun:"
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(allEmojis))

	return emoji.Emoji(allEmojis[randomIndex])
}

var allEmojis = []string{
	":capricorn:",
	":end:",
	":no_mobile_phones:",
	":couple:",
	":snowman:",
	":sunrise_over_mountains:",
	":suspension_railway:",
	":arrows_counterclockwise:",
	":bug:",
	":confused:",
	":dress:",
	":honeybee:",
	":waning_crescent_moon:",
	":balloon:",
	":bus:",
	":package:",
	":pencil2:",
	":rage:",
	":space_invader:",
	":white_medium_small_square:",
	":fast_forward:",
	":rice_cracker:",
	":incoming_envelope:",
	":sa:",
	":womens:",
	":arrow_right:",
	":construction_worker:",
	":notes:",
	":goat:",
	":grey_question:",
	":lantern:",
	":rice_scene:",
	":running:",
	":ferris_wheel:",
	":musical_score:",
	":sparkle:",
	":wink:",
	":art:",
	":clock330:",
	":minidisc:",
	":no_entry_sign:",
	":wind_chime:",
	":cyclone:",
	":herb:",
	":leopard:",
	":banana:",
	":handbag:",
	":honey_pot:",
	":ok:",
	":hearts:",
	":passport_control:",
	":moyai:",
	":smile:",
	":tiger2:",
	":twisted_rightwards_arrows:",
	":children_crossing:",
	":cow:",
	":point_up:",
	":house:",
	":man_with_turban:",
	":mountain_railway:",
	":vibration_mode:",
	":blowfish:",
	":it:",
	":oden:",
	":clock3:",
	":lollipop:",
	":train:",
	":scissors:",
	":triangular_ruler:",
	":wedding:",
	":flashlight:",
	":secret:",
	":sushi:",
	":blue_car:",
	":cd:",
	":milky_way:",
	":mortar_board:",
	":crown:",
	":speech_balloon:",
	":bento:",
	":grey_exclamation:",
	":hotel:",
	":keycap_ten:",
	":newspaper:",
	":outbox_tray:",
	":racehorse:",
	":laughing:",
	":black_large_square:",
	":books:",
	":eight_spoked_asterisk:",
	":heavy_check_mark:",
	":m:",
	":wave:",
	":bicyclist:",
	":cocktail:",
	":european_castle:",
	":point_down:",
	":tokyo_tower:",
	":battery:",
	":dancer:",
	":repeat:",
	":ru:",
	":new_moon:",
	":church:",
	":date:",
	":earth_americas:",
	":footprints:",
	":libra:",
	":mountain_cableway:",
	":small_red_triangle_down:",
	":top:",
	":sunglasses:",
	":abcd:",
	":cl:",
	":ski:",
	":book:",
	":hourglass_flowing_sand:",
	":stuck_out_tongue_closed_eyes:",
	":cold_sweat:",
	":headphones:",
	":confetti_ball:",
	":gemini:",
	":new:",
	":pray:",
	":watch:",
	":coffee:",
	":ghost:",
	":on:",
	":pouch:",
	":taxi:",
	":hocho:",
	":yum:",
	":heavy_plus_sign:",
	":tada:",
	":arrow_heading_down:",
	":clock530:",
	":poultry_leg:",
	":elephant:",
	":gb:",
	":mahjong:",
	":rice:",
	":musical_note:",
	":beginner:",
	":small_red_triangle:",
	":tomato:",
	":clock1130:",
	":japanese_castle:",
	":sun_with_face:",
	":four:",
	":microphone:",
	":tennis:",
	":arrow_up_down:",
	":cn:",
	":horse_racing:",
	":no_bicycles:",
	":snail:",
	":free:",
	":beetle:",
	":black_small_square:",
	":file_folder:",
	":hushed:",
	":skull:",
	":ab:",
	":rocket:",
	":sweet_potato:",
	":guitar:",
	":poodle:",
	":tulip:",
	":large_orange_diamond:",
	":-1:",
	":chart_with_upwards_trend:",
	":de:",
	":grapes:",
	":ideograph_advantage:",
	":japanese_ogre:",
	":telephone:",
	":clock230:",
	":hourglass:",
	":leftwards_arrow_with_hook:",
	":sparkler:",
	":black_joker:",
	":clock730:",
	":first_quarter_moon_with_face:",
	":man:",
	":clock4:",
	":fishing_pole_and_fish:",
	":tophat:",
	":white_medium_square:",
	":mega:",
	":spaghetti:",
	":dart:",
	":girl:",
	":womans_hat:",
	":bullettrain_front:",
	":department_store:",
	":heartbeat:",
	":palm_tree:",
	":swimmer:",
	":yellow_heart:",
	":arrow_upper_right:",
	":clock2:",
	":high_heel:",
	":arrow_double_up:",
	":cry:",
	":dvd:",
	":e-mail:",
	":baby_bottle:",
	":cool:",
	":floppy_disk:",
	":iphone:",
	":minibus:",
	":rooster:",
	":three:",
	":white_small_square:",
	":cancer:",
	":question:",
	":sake:",
	":birthday:",
	":dog2:",
	":loudspeaker:",
	":arrow_up_small:",
	":camel:",
	":koala:",
	":mag_right:",
	":soccer:",
	":bike:",
	":ear_of_rice:",
	":shit:",
	":u7981:",
	":bath:",
	":baby:",
	":lock_with_ink_pen:",
	":necktie:",
	":bikini:",
	":blush:",
	":heartpulse:",
	":pig_nose:",
	":straight_ruler:",
	":u6e80:",
	":gift:",
	":traffic_light:",
	":hibiscus:",
	":couple_with_heart:",
	":pushpin:",
	":u6709:",
	":walking:",
	":grinning:",
	":hash:",
	":radio_button:",
	":raised_hand:",
	":shaved_ice:",
	":barber:",
	":cat:",
	":heavy_exclamation_mark:",
	":ice_cream:",
	":mask:",
	":pig2:",
	":triangular_flag_on_post:",
	":arrow_upper_left:",
	":bee:",
	":beer:",
	":black_nib:",
	":exclamation:",
	":dog:",
	":fire:",
	":ant:",
	":broken_heart:",
	":chart:",
	":clock1:",
	":bomb:",
	":virgo:",
	":a:",
	":fork_and_knife:",
	":copyright:",
	":curly_loop:",
	":full_moon:",
	":shoe:",
	":european_post_office:",
	":ng:",
	":office:",
	":raising_hand:",
	":revolving_hearts:",
	":aquarius:",
	":electric_plug:",
	":meat_on_bone:",
	":mens:",
	":briefcase:",
	":ship:",
	":anchor:",
	":ballot_box_with_check:",
	":bear:",
	":beers:",
	":dromedary_camel:",
	":nut_and_bolt:",
	":construction:",
	":golf:",
	":toilet:",
	":blue_book:",
	":boom:",
	":deciduous_tree:",
	":kissing_closed_eyes:",
	":smiley_cat:",
	":fuelpump:",
	":kiss:",
	":clock10:",
	":sheep:",
	":white_flower:",
	":boar:",
	":currency_exchange:",
	":facepunch:",
	":flower_playing_cards:",
	":person_frowning:",
	":poop:",
	":satisfied:",
	":8ball:",
	":disappointed_relieved:",
	":panda_face:",
	":ticket:",
	":us:",
	":waxing_crescent_moon:",
	":dragon:",
	":gun:",
	":mount_fuji:",
	":new_moon_with_face:",
	":star2:",
	":grimacing:",
	":confounded:",
	":congratulations:",
	":custard:",
	":frowning:",
	":maple_leaf:",
	":police_car:",
	":cloud:",
	":jeans:",
	":fish:",
	":wavy_dash:",
	":clock5:",
	":santa:",
	":japan:",
	":oncoming_taxi:",
	":whale:",
	":arrow_forward:",
	":kissing_heart:",
	":bullettrain_side:",
	":fearful:",
	":moneybag:",
	":runner:",
	":mailbox:",
	":sandal:",
	":zzz:",
	":apple:",
	":arrow_heading_up:",
	":family:",
	":heavy_minus_sign:",
	":saxophone:",
	":u5272:",
	":black_square_button:",
	":bouquet:",
	":love_letter:",
	":metro:",
	":small_blue_diamond:",
	":thought_balloon:",
	":arrow_up:",
	":no_pedestrians:",
	":smirk:",
	":blue_heart:",
	":large_blue_diamond:",
	":vs:",
	":v:",
	":wheelchair:",
	":couplekiss:",
	":tent:",
	":purple_heart:",
	":relaxed:",
	":accept:",
	":green_heart:",
	":pouting_cat:",
	":tram:",
	":bangbang:",
	":collision:",
	":convenience_store:",
	":person_with_blond_hair:",
	":uk:",
	":peach:",
	":tired_face:",
	":bread:",
	":mailbox_closed:",
	":open_mouth:",
	":pig:",
	":put_litter_in_its_place:",
	":u7a7a:",
	":bulb:",
	":clock9:",
	":envelope_with_arrow:",
	":pisces:",
	":baggage_claim:",
	":egg:",
	":sweat_smile:",
	":boat:",
	":fr:",
	":heavy_division_sign:",
	":muscle:",
	":paw_prints:",
	":arrow_left:",
	":black_circle:",
	":kissing_smiling_eyes:",
	":star:",
	":steam_locomotive:",
	":1234:",
	":clock130:",
	":kr:",
	":monorail:",
	":school:",
	":seven:",
	":baby_chick:",
	":bridge_at_night:",
	":hotsprings:",
	":rose:",
	":love_hotel:",
	":princess:",
	":ramen:",
	":scroll:",
	":tropical_fish:",
	":heart_eyes_cat:",
	":information_desk_person:",
	":mouse:",
	":no_smoking:",
	":post_office:",
	":stars:",
	":arrow_double_down:",
	":unlock:",
	":arrow_backward:",
	":hand:",
	":hospital:",
	":ocean:",
	":mountain_bicyclist:",
	":octopus:",
	":sos:",
	":dizzy_face:",
	":tongue:",
	":train2:",
	":checkered_flag:",
	":orange_book:",
	":sound:",
	":aerial_tramway:",
	":bell:",
	":dragon_face:",
	":flipper:",
	":ok_woman:",
	":performing_arts:",
	":postal_horn:",
	":clock1030:",
	":email:",
	":green_book:",
	":point_up_2:",
	":high_brightness:",
	":running_shirt_with_sash:",
	":bookmark:",
	":sob:",
	":arrow_lower_right:",
	":point_left:",
	":purse:",
	":sparkles:",
	":black_medium_small_square:",
	":pound:",
	":rabbit:",
	":woman:",
	":negative_squared_cross_mark:",
	":open_book:",
	":smiling_imp:",
	":spades:",
	":baseball:",
	":fountain:",
	":joy:",
	":lipstick:",
	":partly_sunny:",
	":ram:",
	":red_circle:",
	":cop:",
	":green_apple:",
	":registered:",
	":+1:",
	":crying_cat_face:",
	":innocent:",
	":mobile_phone_off:",
	":underage:",
	":dolphin:",
	":busts_in_silhouette:",
	":umbrella:",
	":angel:",
	":small_orange_diamond:",
	":sunflower:",
	":link:",
	":notebook:",
	":oncoming_bus:",
	":bookmark_tabs:",
	":calendar:",
	":izakaya_lantern:",
	":mans_shoe:",
	":name_badge:",
	":closed_lock_with_key:",
	":fist:",
	":id:",
	":ambulance:",
	":musical_keyboard:",
	":ribbon:",
	":seedling:",
	":tv:",
	":football:",
	":nail_care:",
	":seat:",
	":alarm_clock:",
	":money_with_wings:",
	":relieved:",
	":womans_clothes:",
	":lips:",
	":clubs:",
	":house_with_garden:",
	":sunrise:",
	":monkey:",
	":six:",
	":smiley:",
	":feet:",
	":waning_gibbous_moon:",
	":yen:",
	":baby_symbol:",
	":signal_strength:",
	":boy:",
	":busstop:",
	":computer:",
	":night_with_stars:",
	":older_woman:",
	":parking:",
	":trumpet:",
	":100:",
	":sweat_drops:",
	":wc:",
	":b:",
	":cupid:",
	":five:",
	":part_alternation_mark:",
	":snowboarder:",
	":warning:",
	":white_large_square:",
	":zap:",
	":arrow_down_small:",
	":clock430:",
	":expressionless:",
	":phone:",
	":roller_coaster:",
	":lemon:",
	":one:",
	":christmas_tree:",
	":hankey:",
	":hatched_chick:",
	":u7533:",
	":large_blue_circle:",
	":up:",
	":wine_glass:",
	":x:",
	":nose:",
	":rewind:",
	":two_hearts:",
	":envelope:",
	":oncoming_automobile:",
	":ophiuchus:",
	":ring:",
	":tropical_drink:",
	":turtle:",
	":crescent_moon:",
	":koko:",
	":microscope:",
	":rugby_football:",
	":smoking:",
	":anger:",
	":aries:",
	":city_sunset:",
	":clock1230:",
	":mailbox_with_no_mail:",
	":movie_camera:",
	":pager:",
	":zero:",
	":bank:",
	":eight_pointed_black_star:",
	":knife:",
	":u7121:",
	":customs:",
	":melon:",
	":rowboat:",
	":corn:",
	":eggplant:",
	":heart_decoration:",
	":rotating_light:",
	":round_pushpin:",
	":cat2:",
	":chocolate_bar:",
	":no_bell:",
	":radio:",
	":droplet:",
	":hamburger:",
	":fire_engine:",
	":heart:",
	":potable_water:",
	":telephone_receiver:",
	":dash:",
	":globe_with_meridians:",
	":guardsman:",
	":heavy_multiplication_x:",
	":chart_with_downwards_trend:",
	":imp:",
	":earth_asia:",
	":mouse2:",
	":notebook_with_decorative_cover:",
	":telescope:",
	":trolleybus:",
	":card_index:",
	":euro:",
	":dollar:",
	":fax:",
	":mailbox_with_mail:",
	":raised_hands:",
	":disappointed:",
	":foggy:",
	":person_with_pouting_face:",
	":statue_of_liberty:",
	":dolls:",
	":light_rail:",
	":pencil:",
	":speak_no_evil:",
	":calling:",
	":clock830:",
	":cow2:",
	":hear_no_evil:",
	":scream_cat:",
	":smile_cat:",
	":tractor:",
	":clock11:",
	":doughnut:",
	":hammer:",
	":loop:",
	":moon:",
	":soon:",
	":cinema:",
	":factory:",
	":flushed:",
	":mute:",
	":neutral_face:",
	":scorpius:",
	":wolf:",
	":clapper:",
	":joy_cat:",
	":pensive:",
	":sleeping:",
	":credit_card:",
	":leo:",
	":man_with_gua_pi_mao:",
	":open_hands:",
	":tea:",
	":arrow_down:",
	":nine:",
	":punch:",
	":slot_machine:",
	":clap:",
	":information_source:",
	":tiger:",
	":city_sunrise:",
	":dango:",
	":thumbsdown:",
	":u6307:",
	":curry:",
	":cherries:",
	":clock6:",
	":clock7:",
	":older_man:",
	":oncoming_police_car:",
	":syringe:",
	":heavy_dollar_sign:",
	":open_file_folder:",
	":arrow_right_hook:",
	":articulated_lorry:",
	":dancers:",
	":kissing_cat:",
	":rainbow:",
	":u5408:",
	":boot:",
	":carousel_horse:",
	":fried_shrimp:",
	":lock:",
	":non-potable_water:",
	":o:",
	":persevere:",
	":diamond_shape_with_a_dot_inside:",
	":fallen_leaf:",
	":massage:",
	":volcano:",
	":gem:",
	":shower:",
	":speaker:",
	":last_quarter_moon_with_face:",
	":mag:",
	":anguished:",
	":monkey_face:",
	":sunny:",
	":tangerine:",
	":point_right:",
	":railway_car:",
	":triumph:",
	":two:",
	":gift_heart:",
	":ledger:",
	":sagittarius:",
	":snowflake:",
	":abc:",
	":horse:",
	":ok_hand:",
	":video_camera:",
	":sparkling_heart:",
	":taurus:",
	":frog:",
	":hamster:",
	":helicopter:",
	":fries:",
	":mushroom:",
	":penguin:",
	":truck:",
	":bar_chart:",
	":evergreen_tree:",
	":bow:",
	":clock12:",
	":four_leaf_clover:",
	":inbox_tray:",
	":smirk_cat:",
	":two_men_holding_hands:",
	":water_buffalo:",
	":alien:",
	":video_game:",
	":candy:",
	":page_facing_up:",
	":watermelon:",
	":white_check_mark:",
	":blossom:",
	":crocodile:",
	":no_mouth:",
	":o2:",
	":shirt:",
	":clock8:",
	":eyes:",
	":rabbit2:",
	":tanabata_tree:",
	":wrench:",
	":es:",
	":trophy:",
	":two_women_holding_hands:",
	":clock630:",
	":pineapple:",
	":stuck_out_tongue:",
	":angry:",
	":athletic_shoe:",
	":cookie:",
	":flags:",
	":game_die:",
	":bird:",
	":jack_o_lantern:",
	":ox:",
	":paperclip:",
	":sleepy:",
	":astonished:",
	":back:",
	":closed_book:",
	":hatching_chick:",
	":arrows_clockwise:",
	":car:",
	":ear:",
	":haircut:",
	":icecream:",
	":bust_in_silhouette:",
	":diamonds:",
	":no_good:",
	":pizza:",
	":chicken:",
	":eyeglasses:",
	":see_no_evil:",
	":earth_africa:",
	":fireworks:",
	":page_with_curl:",
	":rice_ball:",
	":white_square_button:",
	":cake:",
	":red_car:",
	":tm:",
	":unamused:",
	":fish_cake:",
	":key:",
	":speedboat:",
	":closed_umbrella:",
	":pear:",
	":satellite:",
	":scream:",
	":first_quarter_moon:",
	":jp:",
	":repeat_one:",
	":shell:",
	":interrobang:",
	":trident:",
	":u55b6:",
	":atm:",
	":door:",
	":kissing:",
	":six_pointed_star:",
	":thumbsup:",
	":u6708:",
	":do_not_litter:",
	":whale2:",
	":school_satchel:",
	":cactus:",
	":clipboard:",
	":dizzy:",
	":waxing_gibbous_moon:",
	":camera:",
	":capital_abcd:",
	":leaves:",
	":left_luggage:",
	":bamboo:",
	":bowling:",
	":eight:",
	":kimono:",
	":left_right_arrow:",
	":stuck_out_tongue_winking_eye:",
	":surfer:",
	":sweat:",
	":violin:",
	":postbox:",
	":bride_with_veil:",
	":recycle:",
	":station:",
	":vhs:",
	":crossed_flags:",
	":memo:",
	":no_entry:",
	":white_circle:",
	":arrow_lower_left:",
	":chestnut:",
	":crystal_ball:",
	":last_quarter_moon:",
	":loud_sound:",
	":strawberry:",
	":worried:",
	":circus_tent:",
	":weary:",
	":bathtub:",
	":snake:",
	":grin:",
	":symbols:",
	":airplane:",
	":heart_eyes:",
	":sailboat:",
	":stew:",
	":tshirt:",
	":rat:",
	":black_medium_square:",
	":clock930:",
	":full_moon_with_face:",
	":japanese_goblin:",
	":restroom:",
	":vertical_traffic_light:",
	":basketball:",
	":cherry_blossom:",
	":low_brightness:",
	":pill:",
}