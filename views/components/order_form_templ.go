// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func OrderForm() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form hx-post=\"/orders\" hx-trigger=\"submit\" hx-swap=\"none\" onsubmit=\"htmx:afterOnLoad: window.location.reload()\" class=\"w-96\"><input type=\"text\" name=\"patient_first_name\" class=\"w-full border rounded-lg mb-2 p-4\" placeholder=\"Patient First Name\" required> <input type=\"text\" name=\"patient_last_name\" class=\"w-full border rounded-lg mb-2 p-4\" placeholder=\"Patient Last Name\" required> <input type=\"text\" name=\"provider_first_name\" class=\"w-full border rounded-lg mb-2 p-4\" placeholder=\"Provider First Name\" required> <input type=\"text\" name=\"provider_last_name\" class=\"w-full border rounded-lg mb-2 p-4\" placeholder=\"Provider Last Name\" required> <button class=\"py-1 px-4 w-full h-10 rounded-lg text-white bg-zinc-800\">Create</button></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
