// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "gothwind/cmd/web/view/components"

func Home() templ.Component {
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
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Hero Section --> <section class=\"bg-blue-600 text-white py-20\"><div class=\"container mx-auto px-4 text-center\"><h1 class=\"text-5xl font-bold\">Welcome to MyWebsite</h1><p class=\"text-lg mt-4\">Your one-stop solution for all your needs.</p><button class=\"mt-6 px-6 py-2 bg-white text-blue-600 font-semibold rounded-md hover:bg-gray-100\">Get Started</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = components.GreetForm().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></section><!-- Content Section --> <section class=\"container mx-auto px-4 py-16\"><div class=\"grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8\"><div class=\"bg-white shadow-lg rounded-lg p-6\"><h2 class=\"text-xl font-bold mb-4\">Feature One</h2><p class=\"text-gray-600\">Description of the first feature goes here.</p></div><div class=\"bg-white shadow-lg rounded-lg p-6\"><h2 class=\"text-xl font-bold mb-4\">Feature Two</h2><p class=\"text-gray-600\">Description of the second feature goes here.</p></div><div class=\"bg-white shadow-lg rounded-lg p-6\"><h2 class=\"text-xl font-bold mb-4\">Feature Three</h2><p class=\"text-gray-600\">Description of the third feature goes here.</p></div></div></section><!-- Content Section --> <section class=\"container mx-auto px-4 py-16\"><div class=\"grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8\"><div class=\"bg-white shadow-lg rounded-lg p-6\"><h2 class=\"text-xl font-bold mb-4\">Feature One</h2><p class=\"text-gray-600\">Description of the first feature goes here.</p></div><div class=\"bg-white shadow-lg rounded-lg p-6\"><h2 class=\"text-xl font-bold mb-4\">Feature Two</h2><p class=\"text-gray-600\">Description of the second feature goes here.</p></div><div class=\"bg-white shadow-lg rounded-lg p-6\"><h2 class=\"text-xl font-bold mb-4\">Feature Three</h2><p class=\"text-gray-600\">Description of the third feature goes here.</p></div></div></section>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Base("Home").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
