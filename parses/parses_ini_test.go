package parses

import "testing"

func TestGetValue(t *testing.T) {
	in := []byte(`
[LIVECOREMODE]
LiveCore_Mode=0

[LIVECORESHOW]
LiveCore_Show_Graphics=0
LiveCore_Show_GraphicsFont=24

[LIVECOREWALLPAPERMODE]
LiveCore_Wallpaper_Mode=0
LiveCore_Wallpaper_Audio=1

[LIVECORELOGMODE]
LiveCore_Log_Process=0

[LIVECOREPLAYMODE]
LiveCore_Play_Mode=2

[LIVECOREWINDOW]
LiveCore_Window_Handle=9506642

[LIVECOREVIDEOADDRESS]
LiveCore_Video_Mode=1
LiveCore_Video_Name=Wallpaper1.pak
LiveCore_Video_Address=C:\Users\10295\Videos\29156949_22_0.flv
`)
	r := getValue(in, "LIVECORESHOW", "LiveCore_Show_Graphics")
	if string(r) != "0" {
		t.Fatal("Error parses ini...")
	}
}

func BenchmarkGetValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := []byte(`
[LIVECOREMODE]
LiveCore_Mode=0

[LIVECORESHOW]
LiveCore_Show_Graphics=0
LiveCore_Show_GraphicsFont=24

[LIVECOREWALLPAPERMODE]
LiveCore_Wallpaper_Mode=0
LiveCore_Wallpaper_Audio=1

[LIVECORELOGMODE]
LiveCore_Log_Process=0

[LIVECOREPLAYMODE]
LiveCore_Play_Mode=2

[LIVECOREWINDOW]
LiveCore_Window_Handle=9506642

[LIVECOREVIDEOADDRESS]
LiveCore_Video_Mode=1
LiveCore_Video_Name=Wallpaper1.pak
LiveCore_Video_Address=C:\Users\10295\Videos\29156949_22_0.flv
`)
		r := getValue(in, "LIVECORESHOW", "LiveCore_Show_Graphics")
		if string(r) != "0" {
			b.Fatal("Error parses ini...")
		}
	}
}
