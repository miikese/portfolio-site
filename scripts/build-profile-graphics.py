#!/usr/bin/env python3
"""Create an original code-drawn motion banner and matching monogram asset.

Needs Pillow. The website uses native HTML/CSS; the GIF is for GitHub.
"""
from pathlib import Path
import math
from PIL import Image, ImageDraw, ImageFont

ROOT = Path(__file__).resolve().parents[1]
FONT_ROOT = Path('/usr/share/fonts/truetype/dejavu')
NAVY, BLUE, IVORY, MUTED, WARM = '#152b4b', '#234fce', '#f8f7f3', '#c4d1e6', '#f0b092'

def font(size, bold=False, mono=False):
    return ImageFont.truetype(str(FONT_ROOT / ('DejaVuSansMono.ttf' if mono else 'DejaVuSans-Bold.ttf' if bold else 'DejaVuSans.ttf')), size)

def banner(phase):
    image = Image.new('RGB', (1200, 400), NAVY)
    d = ImageDraw.Draw(image)
    for x in range(0, 1200, 32): d.line((x, 0, x, 400), fill='#203753')
    for y in range(0, 400, 32): d.line((0, y, 1200, y), fill='#203753')
    d.line((48, 68, 48, 42, 74, 42), fill=WARM, width=2)
    d.line((1126, 358, 1152, 358, 1152, 332), fill=WARM, width=2)
    d.text((70, 68), 'BUILD WITH PURPOSE. STAY CURIOUS.', fill='#b4c8ff', font=font(13, mono=True))
    d.text((68, 108), 'Michael Ikese', fill=IVORY, font=font(54, bold=True))
    d.text((68, 171), 'Emmanuel.', fill=IVORY, font=font(54, bold=True))
    d.text((72, 250), 'Software engineer & programmer', fill=MUTED, font=font(18))
    d.text((72, 282), 'Project manager in training · Cybersecurity enthusiast', fill=MUTED, font=font(14))
    d.text((72, 328), 'NIGERIA / @MIIKESE / OPEN TO OPPORTUNITIES', fill='#b4c8ff', font=font(12, mono=True))
    for radius in (108, 147):
        d.ellipse((947-radius, 192-radius, 947+radius, 192+radius), outline='#57749b', width=1)
    d.line((800, 192, 1094, 192), fill='#344e72')
    d.line((947, 45, 947, 339), fill='#344e72')
    d.rectangle((884, 129, 1010, 255), fill=BLUE, outline='#b4c8ff', width=1)
    d.text((897, 143), 'mi', fill=IVORY, font=font(72, bold=True))
    for angle in (0, 2*math.pi/3, 4*math.pi/3):
        px, py = 947+147*math.cos(angle+phase), 192+147*math.sin(angle+phase)
        d.ellipse((px-4, py-4, px+4, py+4), fill=WARM)
    d.text((822, 360), 'ENGINEERING / LEARNING / DELIVERY', fill='#9fb5d4', font=font(10, mono=True))
    return image

frames = [banner(2*math.pi*i/32).quantize(colors=96) for i in range(32)]
frames[0].save(ROOT/'static/img/profile-banner.gif', save_all=True, append_images=frames[1:], duration=160, loop=0, optimize=True, disposal=1)
image = Image.new('RGB', (512, 512), NAVY)
d = ImageDraw.Draw(image)
for radius in (198,158): d.ellipse((256-radius,256-radius,256+radius,256+radius), outline='#57749b', width=2)
d.line((58,256,454,256), fill='#57749b'); d.line((256,58,256,454), fill='#57749b')
d.rectangle((142,142,370,370), fill=BLUE, outline='#b4c8ff', width=2)
d.text((163,159), 'mi', font=font(140,bold=True), fill=IVORY)
for px,py in [(256,58),(454,256)]: d.ellipse((px-8,py-8,px+8,py+8),fill=WARM)
d.rectangle((95,454,417,500), fill=NAVY)
d.text((256,462), 'MIIKESE', font=font(23,mono=True), fill='#b4c8ff', anchor='mt')
image.save(ROOT/'static/img/profile-mark.png',optimize=True)
print('Built original motion banner and profile monogram.')
