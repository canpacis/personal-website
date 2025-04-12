package app

import (
	"time"

	"github.com/canpacis/pacis/pages"
	"github.com/canpacis/pacis/pages/i18n"
	. "github.com/canpacis/pacis/ui/components"
	. "github.com/canpacis/pacis/ui/html"
	"github.com/canpacis/pacis/ui/icons"
	"github.com/canpacis/personal-website/components"
)

type ReachOutButton struct {
	Label string
	Icon  func(...I) Element
	Href  string
}

type Skill struct {
	Label string
}

type Experience struct {
	Start   time.Time
	End     time.Time
	Company string
	Role    string
}

type ProjectType int

const (
	App = ProjectType(iota)
	Library
	Contribution
)

func (pt ProjectType) String() string {
	switch pt {
	case App:
		return "App"
	case Library:
		return "Library"
	case Contribution:
		return "Contribution"
	default:
		return "Unknown Project Type"
	}
}

type Project struct {
	Type        ProjectType
	Title       string
	Description string
	Link        string
}

func formatTime(date time.Time) string {
	if date.Year() == 0 {
		return "Current"
	}
	return date.Format("Jan 2006")
}

func HomePage(ctx *pages.PageContext) I {
	reach := []ReachOutButton{
		{"X/Twitter", components.XIcon, "https://x.com/can_pacis"},
		{"LinkedIn", components.LinkedInIcon, "https://www.linkedin.com/feed/"},
		{"Github", components.GithubIcon, "https://github.com/canpacis"},
		{"Email", components.EmailIcon, "mailto:canpacis@gmail.com"},
	}

	skills := []Skill{
		{"JavaScript"},
		{"TypeScript"},
		{"Go"},
		{"React"},
		{"Next JS"},
		{"Svelte"},
		{"Rust"},
		{"RestAPI"},
		{"RPC"},
		{"tRPC"},
		{"gRPC"},
		{"PostgreSQL"},
		{"MongoDB"},
		{"Redis"},
		{"AWS"},
		{"GCP"},
		{"Docker"},
		{"Nixpacks"},
	}

	loc := time.Now().Location()

	experience := []Experience{
		{
			time.Date(2025, time.February, 1, 0, 0, 0, 0, loc),
			time.Date(0, time.January, 1, 0, 0, 0, 0, loc),
			"Mitaex",
			"Mobile Engineer",
		},
		{
			time.Date(2024, time.January, 1, 0, 0, 0, 0, loc),
			time.Date(0, time.January, 1, 0, 0, 0, 0, loc),
			"Various",
			"Software Developer",
		},
		{
			time.Date(2022, time.July, 1, 0, 0, 0, 0, loc),
			time.Date(2023, time.August, 1, 0, 0, 0, 0, loc),
			"Macellan",
			"Senior Frontend Developer",
		},
		{
			time.Date(2021, time.February, 1, 0, 0, 0, 0, loc),
			time.Date(2022, time.July, 1, 0, 0, 0, 0, loc),
			"Viavis",
			"Web Developer",
		},
		{
			time.Date(2020, time.January, 1, 0, 0, 0, 0, loc),
			time.Date(2021, time.February, 1, 0, 0, 0, 0, loc),
			"Various",
			"Web Developer",
		},
	}

	projects := []Project{
		{App, "Gaste", "Shorten your links, expand your reach. Gaste is a modern URL shortening and campaign manager/analytics tool", "https://www.gaste.news/"},
		{App, "PratikMaymun", "A free news & magazine outlet with modern web features.", "https://www.pratikmaymun.com/"},
		{App, "Kâşif", "Web based cross platform file explorer. With a modern, slick and exciting interface, Kâşif makes file browsing easy and fun.", "https://github.com/kasif-apps/"},
		{App, "Clipboard", "A MacOS application that creates a clipboard on your tray. You can copy text, images and files and they will show up in your tray!", "https://github.com/canpacis/Clipboard"},
		{Library, "Pacis", "An SSR UI library for Go with Tailwind for styling and Alpine.js for interactivity.", "https://ui.canpacis.com/"},
		{Library, "Moonbite", "Moonbite is an experimental type safe language that is compiled to custom byte code.", "https://github.com/moonbite-org/moonbite"},
		{Library, "Affixi", "A helper library for Turkish noun suffixes written in typescript.", "https://github.com/canpacis/affixi"},
		{Library, "Casbin Adapter", "A simple casbin adapter implementation that uses go's pgx driver. It implement filter, context and batch interfaces of casbin adapters.", "https://github.com/canpacis/casbin-pgx-context-adapter"},
		{Library, "Brainfuck Interpreter", "A brainfuck interpreter that is built with a lexer, a parser and a debugger protocol written in Go. This is actually intended to be a superset of brainfuck so there are extended capabilities of the runtime.", "https://github.com/canpacis/brainfuck-interpreter"},
		{Library, "Vedia", "Vedia is a stack based language where you don't have any variables, any meaningful variables that is. Expressions can still yield a value and primitive literals do exist.", "https://github.com/canpacis/vedia"},
		{Contribution, "Mantine", "Build fully functional accessible web applications faster than ever – Mantine includes more than 100 customizable components and 50 hooks to cover you in any situation", "https://mantine.dev/"},
	}

	return Frag(
		Header(
			Class("flex gap-4 justify-between items-center py-12 app-container"),

			Div(
				H1(
					Class("text-2xl mb-2 md:mb-0 md:text-3xl font-bold leading-normal md:leading-14"),

					i18n.Text("heading"),
				),
				P(
					Class("font-light max-w-none md:max-w-[75%] leading-5"),

					i18n.Text("subtitle"),
				),
			),
			Img(
				Src("/public/profile.webp"),
				Class("rounded-full border-4 border-accent/50"),
				Alt("Profile"),
				Width("82"),
				Height("82"),
			),
		),

		Main(
			Class("flex flex-col gap-8 app-container flex-1"),

			Section(
				H2(Class("font-bold mb-2"), i18n.Text("about.title")),
				P(
					Class("text-muted-foreground"),

					i18n.Text("about.text"),
				),
			),

			Section(
				H2(Class("font-bold mb-2"), i18n.Text("reach_out")),
				Div(
					Class("flex gap-2 flex-wrap"),

					Map(reach, func(button ReachOutButton, i int) Node {
						return Button(
							Replace(Knot),
							Class("size-16"),
							Href(button.Href),
							ButtonSizeIcon,
							ButtonVariantGhost,

							button.Icon(Class("size-6")),
							Span(Class("sr-only"), Text(button.Label)),
						)
					}),
				),
			),

			Section(
				H2(Class("font-bold mb-2"), i18n.Text("skills.title")),

				Div(
					Class("flex gap-2 flex-wrap"),

					Map(skills, func(skill Skill, i int) Node {
						return Badge(Text(skill.Label))
					}),
				),
			),

			Section(
				Class("relative"),

				H2(Class("font-bold mb-2"), i18n.Text("experience.title")),
				Ul(
					Class("flex flex-col gap-2 pl-2 before:content-[''] before:h-[90%] before:w-px before:bg-gradient-to-b before:from-transparent before:via-sky-500 before:to-transparent before:block before:absolute before:top-[10%] before:left-0"),

					Map(experience, func(exp Experience, i int) Node {
						return Li(
							Class("w-full hover:bg-accent/50 px-3 py-2 rounded-md"),

							P(Class("font-light leading-snug"), Text(exp.Role)),
							P(
								Class("text-muted-foreground leading-snug"),

								Span(Text(exp.Company)),
								Span(Class("text-xs ml-4"), Text(formatTime(exp.Start)), Span(Class("mx-1"), Text("-")), Text(formatTime(exp.End))),
							),
						)
					}),
				),
			),

			Section(
				H2(Class("font-bold mb-2"), i18n.Text("projects.title")),

				Div(
					Class("grid grid-cols-1 md:grid-cols-2 gap-2"),

					Map(projects, func(project Project, i int) Node {
						return Div(
							Class("rounded-md p-4 h-72 flex flex-col justify-center relative group overflow-hidden project"),

							Div(
								Data("laser", ""),
								Class("absolute border border-transparent w-full h-72 pointer-events-none inset-0 rounded-[inherit] [mask-clip:padding-box,border-box] [mask-composite:intersect] [mask-image:linear-gradient(transparent,transparent),linear-gradient(#000,#000)] opacity-0 group-hover:opacity-100 transition-all"),

								Div(
									Data("offset", ""),
									Class("absolute inset-0 bg-gradient-to-l from-transparent via-purple-500 to-transparent w-72 aspect-square"),
									StyleAttr("offset-path: rect(0px auto auto 0px round 288px)"),
								),
							),

							Div(
								Data("laser", ""),
								Class("absolute border border-transparent w-full h-72 pointer-events-none inset-0 rounded-[inherit] [mask-clip:padding-box,border-box] [mask-composite:intersect] [mask-image:linear-gradient(transparent,transparent),linear-gradient(#000,#000)] opacity-0 group-hover:opacity-100 transition-all"),

								Div(
									Data("offset-offset", ""),
									Class("absolute inset-0 bg-gradient-to-l from-transparent via-sky-500 to-transparent w-72 aspect-square"),
									StyleAttr("offset-path: rect(0px auto auto 0px round 288px)"),
								),
							),

							Badge(BadgeVariantSecondary, Text(project.Type.String())),
							Div(
								Class("flex justify-between"),

								H3(Class("text-base md:text-lg font-bold my-2"), Text(project.Title)),
								Button(
									ButtonSizeIcon,
									ButtonVariantGhost,
									Replace(Knot),
									Href(project.Link),

									icons.ExternalLink(),
									Span(Class("sr-only"), Text("Open")),
								),
							),
							P(
								Class("text-sm md:text-base"),

								Text(project.Description),
							),
						)
					}),
				),
			),

			Section(
				Class("my-10 flex justify-center items-center"),

				Button(
					Replace(Knot),
					Href("https://buymeacoffee.com/canpacis"),
					Class("!rounded-full"),
					ButtonSizeLg,

					i18n.Text("misc.cta"),
				),
			),
		),

		Footer(
			Class("border-t h-10 w-full mt-4"),

			Div(
				Class("app-container flex items-center justify-center text-sm h-full text-muted-foreground"),

				i18n.Text("misc.footer.0"),
				Knot(Class("hover:text-white hover:underline mx-1"), Href("https://github.com/canpacis"), i18n.Text("misc.footer.1")),
				i18n.Text("misc.footer.2"),
				Knot(Class("hover:text-white hover:underline mx-1"), Href("https://github.com/canpacis/pacis"), i18n.Text("misc.footer.3")),
			),
		),
	)
}
