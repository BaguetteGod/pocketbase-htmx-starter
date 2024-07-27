// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package sidebar

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"date-rate/app/components/svgs"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"log"
	"os"
)

func SidebarLayout(c echo.Context) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"{ open: false }\"><!-- Off-canvas menu for mobile, show/hide based on off-canvas menu state. --><div x-show=\"open\" class=\"relative z-50 lg:hidden\" role=\"dialog\" aria-modal=\"true\"><!-- Off-canvas menu backdrop, show/hide based on off-canvas menu state. --><div x-show=\"open\" x-transition:enter=\"transition-opacity ease-linear duration-300\" x-transition:enter-start=\"opacity-0\" x-transition:enter-end=\"opacity-100\" x-transition:leave=\"transition-opacity ease-linear duration-300\" x-transition:leave-start=\"opacity-100\" x-transition:leave-end=\"opacity-0\" class=\"fixed inset-0 bg-gray-900/80\" aria-hidden=\"true\"></div><div class=\"fixed inset-0 flex\"><!-- Off-canvas menu, show/hide based on off-canvas menu state. --><div x-show=\"open\" x-transition:enter=\"transition ease-in-out duration-300 transform\" x-transition:enter-start=\"-translate-x-full\" x-transition:enter-end=\"translate-x-0\" x-transition:leave=\"transition ease-in-out duration-300 transform\" x-transition:leave-start=\"translate-x-0\" x-transition:leave-end=\"-translate-x-full\" class=\"relative mr-16 flex w-full max-w-xs flex-1\"><!-- Close button, show/hide based on off-canvas menu state. --><div x-show=\"open\" x-transition:enter=\"ease-in-out duration-300\" x-transition:enter-start=\"opacity-0\" x-transition:enter-end=\"opacity-100\" x-transition:leave=\"ease-in-out duration-300\" x-transition:leave-start=\"opacity-100\" x-transition:leave-end=\"opacity-0\" class=\"absolute left-full top-0 flex w-16 justify-center pt-5\"><button x-on:click=\"open = !open\" type=\"button\" class=\"-m-2.5 p-2.5\"><span class=\"sr-only\">Close sidebar</span> <svg class=\"h-6 w-6 text-white\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" aria-hidden=\"true\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M6 18L18 6M6 6l12 12\"></path></svg></button></div><!-- Sidebar component, swap this element with another sidebar if you like -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = sidebarContent(c).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div><!-- Static sidebar for desktop --><div class=\"hidden lg:fixed lg:inset-y-0 lg:z-50 lg:flex lg:w-72 lg:flex-col\"><!-- Sidebar component, swap this element with another sidebar if you like -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = sidebarContent(c).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"sticky top-0 z-40 flex items-center gap-x-6 bg-white px-4 py-4 shadow-sm sm:px-6 lg:hidden\"><button x-on:click=\"open = !open\" type=\"button\" class=\"-m-2.5 p-2.5 text-gray-700 lg:hidden\"><span class=\"sr-only\">Open sidebar</span> <svg class=\"h-6 w-6\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" aria-hidden=\"true\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5\"></path></svg></button><div class=\"flex-1 text-sm font-semibold leading-6 text-gray-900\">Dashboard</div><a href=\"#\"><span class=\"sr-only\">Your profile</span> <img class=\"h-8 w-8 rounded-full bg-gray-50\" src=\"https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&amp;ixid=eyJhcHBfaWQiOjEyMDd9&amp;auto=format&amp;fit=facearea&amp;facepad=2&amp;w=256&amp;h=256&amp;q=80\" alt=\"\"></a></div><main class=\"py-10 lg:pl-72\"><div class=\"px-4 sm:px-6 lg:px-8\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></main></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

type link struct {
	href     string
	linkName string
	icon     templ.Component
}

func (l *link) active(c echo.Context) bool {
	url := c.Request().URL
	return url.String() == l.href
}

var links = []link{
	{"/dashboard", "Dashboard", svgs.Home("h-6 w-6 shrink-0 group-hover:text-blue-600")},
	{"/lists", "Lists", svgs.ListBullet("h-6 w-6 shrink-0 group-hover:text-blue-600")},
}

func getUsername(c echo.Context) string {
	return apis.RequestInfo(c).AuthRecord.GetString("username")
}

func getAvatar(c echo.Context) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env")
	}

	appUrl := os.Getenv("APP_URL")
	user := apis.RequestInfo(c).AuthRecord
	filename := user.GetString("avatar")
	avatar := fmt.Sprintf("%s/api/files/%s/%s/%s", appUrl, user.Collection().Id, user.Id, filename)
	return avatar
}

func sidebarContent(c echo.Context) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex grow flex-col gap-y-5 overflow-y-auto bg-white px-6 pb-2 lg:pb-0 border-r-0 lg:border-r lg:border-gray-200\"><div class=\"flex h-16 shrink-0 items-center\"><img class=\"h-8 w-auto\" src=\"https://tailwindui.com/img/logos/mark.svg?color=blue&amp;shade=600\" alt=\"Your Company\"></div><nav class=\"flex flex-1 flex-col\"><ul role=\"list\" class=\"flex flex-1 flex-col gap-y-7\"><li><ul role=\"list\" class=\"-mx-2 space-y-1\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, item := range links {
			templ_7745c5c3_Err = sidebarLink(item, c).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul></li><li><div class=\"text-xs font-semibold leading-6 text-gray-400\">Recent lists</div><ul role=\"list\" class=\"-mx-2 mt-2 space-y-1\"><li><a href=\"#\" class=\"group flex gap-x-3 rounded-md p-2 text-sm font-semibold leading-6 text-gray-700 hover:bg-gray-50 hover:text-blue-600\"><span class=\"flex h-6 w-6 shrink-0 items-center justify-center rounded-lg border border-gray-200 bg-white text-[0.625rem] font-medium text-gray-400 group-hover:border-blue-600 group-hover:text-blue-600\">H</span> <span class=\"truncate\">Heroicons</span></a></li><li><a href=\"#\" class=\"group flex gap-x-3 rounded-md p-2 text-sm font-semibold leading-6 text-gray-700 hover:bg-gray-50 hover:text-blue-600\"><span class=\"flex h-6 w-6 shrink-0 items-center justify-center rounded-lg border border-gray-200 bg-white text-[0.625rem] font-medium text-gray-400 group-hover:border-blue-600 group-hover:text-blue-600\">T</span> <span class=\"truncate\">Tailwind Labs</span></a></li><li><a href=\"#\" class=\"group flex gap-x-3 rounded-md p-2 text-sm font-semibold leading-6 text-gray-700 hover:bg-gray-50 hover:text-blue-600\"><span class=\"flex h-6 w-6 shrink-0 items-center justify-center rounded-lg border border-gray-200 bg-white text-[0.625rem] font-medium text-gray-400 group-hover:border-blue-600 group-hover:text-blue-600\">W</span> <span class=\"truncate\">Workcation</span></a></li></ul></li><li class=\"-mx-6 mt-auto hidden lg:block\"><a href=\"#\" class=\"flex items-center gap-x-4 px-6 py-3 text-sm font-semibold leading-6 text-gray-900 hover:bg-gray-50\"><img class=\"h-8 w-8 rounded-full bg-gray-50 object-cover\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(getAvatar(c))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/layouts/sidebar/sidebar.templ`, Line: 204, Col: 82}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"\"> <span class=\"sr-only\">Your profile</span> <span aria-hidden=\"true\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(getUsername(c))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/layouts/sidebar/sidebar.templ`, Line: 206, Col: 47}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></a></li></ul></nav></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func sidebarLink(link link, c echo.Context) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 = []any{"flex gap-3 leading-6 font-medium p-2 rounded hover:bg-gray-50 group hover:text-blue-600",
			templ.KV("text-blue-600 bg-gray-50", link.active(c)), templ.KV("text-gray-700 [&>svg]:text-gray-400",
				!link.active(c))}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var6...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 templ.SafeURL = templ.SafeURL(link.href)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var7)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var6).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/layouts/sidebar/sidebar.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = link.icon.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(link.linkName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/layouts/sidebar/sidebar.templ`, Line: 223, Col: 18}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></li>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
