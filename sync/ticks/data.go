package main

import (
	"fmt"
	"math/rand"
)

// var messages = []string{
// 	"Salut tout le monde !",
// 	"Hype dans le chat !",
// 	"Quel est le jeu du jour ?",
// 	"PogChamp xD !",
// 	"GG !",
// 	"Quand est le prochain stream ?",
// 	"Qui est votre streamer préféré ?",
// 	"Faites un shoutout à Hachitaka !",
// 	"Quelqu'un a-t-il des conseils pour ce niveau ?",
// 	"C'est mon premier stream ici, bienvenue à moi ! :p",
// 	"LUL, ce fail était épique !",
// 	"Où sont mes modérateurs ?",
// 	"Chat, vous êtes trop bruyant aujourd'hui !",
// 	"Quelqu'un d'autre a déjà joué à ce jeu ?",
// 	"Faites un 1 dans le chat si vous êtes hypé !",
// 	"Les emotes sont tellement drôles ici :p",
// 	"Quelle est votre musique préférée en ce moment ?",
// 	"Bonjour du Maroc !",
// 	"Quand est-ce que vous faites un meet-up ?",
// 	"Je suis nouveau ici, que se passe-t-il ?",
// 	"Quel est le thème du stream aujourd'hui ?",
// 	"Top 5 des jeux de tous les temps, allez !",
// 	"Quels sont vos emotes préférés sur Twitch ?",
// 	"À quelle heure commence le stream demain ?",
// 	"Faites un V dans le chat pour la victoire !",
// 	"GLHF, tout le monde !",
// 	"Quel est votre émote préférée sur Twitch ?",
// 	"Nous atteignons bientôt 100k followers, c'est incroyable !",
// 	"Merci pour le don, Hachitaka !! <3 <3",
// 	"Je cherche des coéquipiers pour un squad, qui est partant ?",
// 	"Des suggestions pour de nouveaux emotes ?",
// 	"Je suis en train de découvrir un nouveau jeu, venez avec moi !",
// 	"Ce moment où le chat est plus divertissant que le jeu lui-même xD",
// 	"Quelqu'un a des conseils pour améliorer la qualité du stream ?",
// 	"N'oubliez pas de suivre et d'activer les notifications !",
// 	"Nous venons d'atteindre un nouveau sub goal, merci à tous !",
// 	"Je suis en train de manger devant l'écran, quel est votre plat préféré ?",
// 	"Shoutout à mes modos, vous faites un super boulot !",
// 	"bienvenue à tous !",
// 	"J'organise un giveaway à la fin du stream, restez jusqu'au bout !",
// 	"C'était un stream épique aujourd'hui, merci d'avoir été là !",
// 	"Qui est hypé pour la prochaine mise à jour du jeu ?",
// 	"Dites-moi, quel est votre moment préféré sur le stream jusqu'à présent ?",
// }

var pseudos = []string{"Hachitaka86", "Dedeschunter", "Bloodgast", "Eugenepoubel", "Cildrim", "Riiaa"}

var messages = []string{
	"Salut les chatons, c'est parti pour une nouvelle aventure avec %s ! 🚀",
	"Hey %s, tu es prêt pour du fun aujourd'hui ? 😎",
	"%s, tu es vraiment le MVP du chat, toujours là pour mettre de l'ambiance ! 🌟",
	"Les amis, on a un nouveau follower dans la famille ! Bienvenue à %s ! 🎉",
	"%s, tu as déjà essayé de jouer à un jeu avec une main et de taper dans le chat avec l'autre ? C'est tout un défi ! 😂",
	"Shoutout à %s, le maître incontesté des émotes dans le chat ! 💯",
	"Qui pense que %s va décrocher le jackpot aujourd'hui ? 🎰",
	"Alerte spoil : %s est sur le point de découvrir un plot twist épique dans le jeu ! 🕵️‍♂️",
	"Les amis, préparez-vous à rire, %s vient de lancer une blague dans le chat ! 😂",
	"%s, on a besoin de ton expertise stratégique ici, quelle est la meilleure approche pour ce boss ? 🤔",
	"%s, tu es le détective ultime du chat, toujours là pour résoudre les mystères du jeu ! 🔍",
	"Est-ce que %s peut battre son propre record de dons aujourd'hui ? Faites monter la cagnotte, les généreux donateurs ! 💸",
	"%s, quel est ton snack préféré pendant les streams ? Peut-être que tu devrais lancer une émission culinaire ! 🍕",
	"Alerte ! %s vient de déclencher un combo émojis dans le chat, c'est une œuvre d'art ! 🎨",
	"%s, tu es le gardien fidèle du chat, toujours prêt à modérer avec style ! ⚔️",
	"Un petit défi pour %s : essaye de prononcer 'supercalifragilisticexpialidocious' en moins de 10 secondes ! 🤪",
	"Salut %s, si tu étais un personnage de jeu, quel serait ton pouvoir spécial ? 💪",
	"Les spectateurs, donnez une note de 1 à 10 à l'entrée en scène de %s dans le chat ! 🌟",
	"%s, on a entendu dire que tu es le pro des Easter eggs. Quel est le plus mémorable que tu aies découvert récemment ? 🥚",
	"Quand %s entre dans le chat, c'est comme si une bouffée d'énergie positive nous enveloppait tous ! ✨",
	"%s, quel serait le titre du film basé sur ta dernière partie ? 🎬",
	"Les viewers, quel est le meilleur surnom que vous donneriez à %s ? Faites preuve de créativité ! 😄",
	// 20 nouvelles phrases
	"%s, si tu étais un super-héros, quel serait ton nom de code ? 🦸‍♂️",
	"Les chatons, préparez-vous à être époustouflés, %s a une nouvelle astuce à nous montrer ! 🎩✨",
	"%s, tu es la force tranquille du chat, toujours là pour apaiser les esprits agités. 🌈",
	"Alerte câlin virtuel pour %s ! 🤗 Qui veut se joindre à la vague de tendresse ? 💖",
	"Les amis, devinez quel jeu %s va massacrer aujourd'hui ! Préparez-vous pour des skills légendaires ! 🔥",
	"Shoutout à %s pour son avatar stylé. Où peut-on acheter une copie ? 👀",
	"%s, tu es notre source d'inspiration quotidienne. Quel est ton secret pour garder cette énergie positive ? 🌞",
	"Les spectateurs, proposez des défis farfelus à %s et voyons s'il les relève en live ! 🤣",
	"Quelle serait la signature de %s s'il devenait une célébrité du gaming ? 📸",
	"%s, si tu pouvais inviter trois personnages de jeu à dîner, qui choisirais-tu ? 🍽️",
	"Alerte aux paparazzis ! %s est repéré en train de streamer dans une nouvelle chemise funky ! 👕",
	"Les viewers, partagez votre théorie la plus farfelue sur le prochain scénario que %s va rencontrer ! 🤔",
	"%s, si tu devais composer une chanson sur ton expérience de jeu, quel serait le titre ? 🎵",
	"Salut %s, quel est ton jeu rétro préféré et pourquoi ? 🕹️",
	"Les chatons, décrivez %s en un seul mot ! Les réponses les plus créatives gagnent des points bonus ! 🏆",
	"%s, on a entendu dire que tu as une réserve secrète de snacks. Peux-tu nous dévoiler ton top 3 ? 🍿",
	"Alerte spoiler inversé : %s va-t-il réussir à éviter ce piège sournois dans le jeu ? Suspens insoutenable ! 😱",
	"%s, si tu pouvais interchanger les rôles avec un personnage de jeu, lequel serait-ce et pourquoi ? 🔄",
	"Les viewers, envoyez des encouragements à %s pour son prochain boss fight ! 💪",
	"%s, tu es tellement OP que même les PNJ t'envient. Comment fais-tu pour être aussi badass ? 🔥",
}

var usernames = []string{
	"ShadowNinja21",
	"PixelPirate",
	"ThunderByte",
	"StarGazer87",
	"LunaCraze",
	"GameKnight42",
	"NeonNebula",
	"SilverBulletX",
	"QuantumSnipe",
	"MysticSpectre",
	"EchoStrike",
	"BlazeFury",
	"FrostyVortex",
	"CyberWraith",
	"RaptorClaw",
	"PhantomStriker",
	"NovaBlast",
	"ZenithSpecter",
	"VelocityVibe",
	"SolarStorm77",
	"TurboTwist",
	"CelestialWiz",
	"PixelPioneer",
	"VortexVagrant",
	"LunarLancer",
}

// Noir: \x1b[30m
// Rouge: \x1b[31m
// Vert: \x1b[32m
// Jaune: \x1b[33m
// Bleu: \x1b[34m
// Magenta (violet): \x1b[35m
// Cyan (bleu clair): \x1b[36m
// Blanc: \x1b[37m
// Réinitialiser (par défaut): \x1b[39m
// bold \x1b[1m
// reset bold \x1b[0m
func fmtUsername(username string) string {
	color := rand.Intn(7-1) + 1
	return fmt.Sprintf("\x1b[1m\x1b[3%vm%v\x1b[39m\x1b[0m", color, username)
}

func format(text, color string) string {
	var colors = map[string]string{
		"black":   "\x1b[30m",
		"red":     "\x1b[31m",
		"green":   "\x1b[32m",
		"yellow":  "\x1b[33m",
		"blue":    "\x1b[34m",
		"magenta": "\x1b[35m",
		"cyan":    "\x1b[36m",
		"white":   "\x1b[37m",
		"bold":    "\x1b[1m",
		"reset":   "\x1b[0m\x1b[39m",
	}
	return fmt.Sprintf("%v%v%v", colors[color], text, colors["reset"])
}

func generateMessage(hour string) string {
	username := usernames[rand.Intn(len(usernames)-1)]
	pseudo := format(pseudos[rand.Intn(len(pseudos)-1)], "bold")
	message := messages[rand.Intn(len(messages)-1)]
	msg := fmt.Sprintf("%v %v: %v", hour, fmtUsername(username), fmt.Sprintf(message, pseudo))
	return msg
}
