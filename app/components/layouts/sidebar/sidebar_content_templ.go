// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package sidebar

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"date-rate/app/components/svgs"
	"date-rate/app/lib"
	"github.com/labstack/echo/v5"
)

var links = []link{
	{"/dashboard", "Dashboard", svgs.Home("h-6 w-6 shrink-0 group-hover:text-blue-600")},
	{"/lists", "Lists", svgs.ListBullet("h-6 w-6 shrink-0 group-hover:text-blue-600")},
}

var linksRecent = []linkRecent{
	{"/dashboard", "Dashboard"},
	{"/lists", "Lists"},
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
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul></li><li><div class=\"text-xs font-semibold leading-6 text-gray-400\">Recent lists</div><ul role=\"list\" class=\"-mx-2 mt-2 space-y-1\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, item := range linksRecent {
			templ_7745c5c3_Err = sidebarRecentLink(item, c).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul></li><li><ul role=\"list\" class=\"-mx-2 space-y-1\"><li><form method=\"post\" action=\"/logout\" hx-boost=\"true\" class=\"mb-0\"><button class=\"w-full flex gap-3 leading-6 font-medium p-2 rounded hover:bg-gray-50 group hover:text-blue-600 text-sm\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = svgs.Logout("h-6 w-6 text-gray-400 group-hover:text-blue-600").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("Logout</button></form></li></ul></li><li class=\"-mx-6 mt-auto hidden lg:flex flex-col gap-0\"><a href=\"/profile\" class=\"w-full flex items-center gap-x-4 px-6 py-3 text-sm font-semibold leading-6 text-gray-900 hover:bg-gray-50\"><img class=\"h-8 w-8 rounded-full bg-gray-50 object-cover\" src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(lib.GetAvatar(c))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/layouts/sidebar/sidebar_content.templ`, Line: 66, Col: 86}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"\"> <span class=\"sr-only\">Your profile</span> <span aria-hidden=\"true\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(lib.GetUsername(c))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/components/layouts/sidebar/sidebar_content.templ`, Line: 68, Col: 51}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
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