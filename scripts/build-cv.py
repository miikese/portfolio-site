#!/usr/bin/env python3
"""Build a selectable-text, single-page CV using only Python's standard library."""
import json
from pathlib import Path
import textwrap

ROOT = Path(__file__).resolve().parents[1]
data = json.loads((ROOT / "cv/content.json").read_text())
commands = []
y = 793
INK = "0.09 0.15 0.22"
MUTED = "0.31 0.37 0.43"
BLUE = "0.14 0.31 0.81"


def literal(value):
    raw = value.encode("cp1252")
    return "(" + "".join(f"\\{b:03o}" if b > 126 or b < 32 else "\\" + chr(b) if chr(b) in "()\\" else chr(b) for b in raw) + ")"


def line(value, size=9.4, bold=False, color=INK, x=44, leading=13):
    global y
    font = "F2" if bold else "F1"
    commands.append(f"BT /{font} {size} Tf {color} rg 1 0 0 1 {x} {y:.1f} Tm {literal(value)} Tj ET")
    y -= leading


def paragraph(value, width=104, **kwargs):
    for part in textwrap.wrap(value, width=width):
        line(part, **kwargs)


def section(label):
    global y
    y -= 12
    line(label.upper(), size=9.3, bold=True, color=BLUE, leading=12)
    commands.append(f"{BLUE} RG 0.5 w 44 {y+2:.1f} m 551 {y+2:.1f} l S")
    y -= 9


line(data["name"], size=24, bold=True, leading=31)
line(data["role"], size=10.2, color=BLUE, leading=20)
line(data["contact"], size=8.8, color=MUTED, leading=14)
line(data["location"], size=8.4, color=MUTED, leading=14)
section("Profile")
paragraph(data["summary"])
section("Skills & development")
for label, value in data["skills"]:
    paragraph(label + ": " + value, size=9.1, leading=13)
section("Selected projects")
for project in data["projects"]:
    line(project["title"], size=10.5, bold=True, leading=15)
    paragraph(project["stack"], size=8.0, color=BLUE, leading=12)
    paragraph(project["description"], size=9.1, leading=13)
    y -= 8
section("Education & ongoing learning")
line(data["education"], size=9.4, bold=True, leading=15)
paragraph(data["education_detail"], size=9.1, leading=13)
y -= 9
paragraph(data["availability"], size=8.7, color=MUTED, leading=12)
if y < 44:
    raise SystemExit(f"CV overflows its page (remaining y={y}); shorten the content.")

stream = "\n".join(commands).encode("ascii")
objects = [
    b"<< /Type /Catalog /Pages 2 0 R >>",
    b"<< /Type /Pages /Kids [3 0 R] /Count 1 >>",
    b"<< /Type /Page /Parent 2 0 R /MediaBox [0 0 595 842] /Resources << /Font << /F1 4 0 R /F2 5 0 R >> >> /Contents 6 0 R /Annots [8 0 R 9 0 R] >>",
    b"<< /Type /Font /Subtype /Type1 /BaseFont /Helvetica /Encoding /WinAnsiEncoding >>",
    b"<< /Type /Font /Subtype /Type1 /BaseFont /Helvetica-Bold /Encoding /WinAnsiEncoding >>",
    f"<< /Length {len(stream)} >>\nstream\n".encode() + stream + b"\nendstream",
    f"<< /Title {literal(data['name'] + ' - CV')} /Author {literal(data['name'])} /Subject (Software engineering, programming, and project management training) >>".encode(),
    b"<< /Type /Annot /Subtype /Link /Rect [44 741 239 756] /Border [0 0 0] /A << /S /URI /URI (mailto:emmanlemichel2019@gmail.com) >> >>",
    b"<< /Type /Annot /Subtype /Link /Rect [337 741 462 756] /Border [0 0 0] /A << /S /URI /URI (https://github.com/miikese) >> >>",
]
result = bytearray(b"%PDF-1.4\n%\xe2\xe3\xcf\xd3\n")
offsets = [0]
for i, obj in enumerate(objects, 1):
    offsets.append(len(result))
    result.extend(f"{i} 0 obj\n".encode() + obj + b"\nendobj\n")
xref = len(result)
result.extend(f"xref\n0 {len(objects)+1}\n0000000000 65535 f \n".encode())
for offset in offsets[1:]:
    result.extend(f"{offset:010d} 00000 n \n".encode())
result.extend(f"trailer\n<< /Size {len(objects)+1} /Root 1 0 R /Info 7 0 R >>\nstartxref\n{xref}\n%%EOF\n".encode())
path = ROOT / "static/cv/resume.pdf"
path.write_bytes(result)
print(f"Built {path} ({len(result)} bytes; selectable text, one A4 page)")
