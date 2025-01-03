external langs = [
    ["Python", 0.5, #3572A5],
    ["Rust",   0.3, #dea584],
    ["JavaScript", 0.2, #f1e05a],
]
external stroke = 1.2
external pad = 15
external bgcolor = #1e1e1e
external outline = #00d9ff
external textcolor = #ffffffc4
external titlecolor = #ffffff

let w, h = @Dimensions()

@DrawRoundedRectangle(stroke / 2, stroke / 2 , w - stroke, h - stroke, 5)

@SetColor(bgcolor)
@FillPreserve()

@SetStrokeWidth(stroke)
@SetColor(outline)
@Stroke()

let opacity = @Tween(0.0, 1.0, "cubic")

@SetColor(@alpha(titlecolor, opacity))
@SetFontSize(30)
@SetFont("ggsans-bold")

let t = "Most Used Languages"
t = @truncate(t, @Tween(1, @len(t) + 5, "cubic_out"), "")
@DrawString(t, pad, pad - 5)
@Fill()



with @PushRelative(pad, pad + 30, w - pad * 2, h) {
    let mw = @Width()

    @DrawRoundedRectangle(0, 0, mw, 15, 2)
    @Clip()

    let offset = 0
    for i in 0 : @len(langs) {
        let _, per, color = @get(langs, i)
        let bw = @Tween(0.0, mw * per, "cubic")

        @DrawRectangle(offset, 0, bw, 15)
        @SetColor(@alpha(color, opacity))
        @Fill()    

        offset = offset + bw
    }

    with @PushRelative(10, 35, mw - 20, h) {
        @SetFont("ggsans-semibold")
        @SetFontSize(18)

        let swap_at = @len(langs) / 2
        if @len(langs) % 2 == 0 {
            swap_at = swap_at - 1
        }

        let offsetx = 0
        let offsety = 0
        let ax = 0.0
        let o = 12
        let swapped = false

        for i in 0 : @len(langs) {
            let name, per, color = langs[i]

            let base = @float(@len(langs) / 1.0)
            let a = if swapped {
                @TweenEx(0.0, 1.0, (i - @float(swap_at)) / base, 1.0, "linear")
            } else {
                @TweenEx(0.0, 1.0, i / base, 1.0, "linear")
            }

            // let a = 

            @SetColor(@alpha(color, a))
            @DrawCircle(offsetx, offsety, 6)
            @Fill()

            @SetColor(@alpha(textcolor, a))
            name = @replace(name, "CodeQL", "Quilt")

            per = @Tween(per, per, "cubic") * 100
            @DrawStringAnchored(name .. " " .. @fix(per, 1) .. "%", offsetx + o, offsety - 9, ax, 0.0)
            @Fill()

            offsety = offsety + 25

            if i == swap_at { 
                offsety = 0
                offsetx = mw - 20
                o = -12
                ax = 1.0
                swapped = true
            }
        }
    }
}

// let blur = 0.0
// if blur > 0.0 {
    // let img = @Screenshot()
    // @BlurInplace(img, blur)
    // @DrawImage(img, 0, 0)
    // @DeleteImage(img)
// }
