from pathlib import Path


def define_visitor(program: list[str], base_name, types: dict[str, list[str]]):
  program.append("type Visitor interface {")
  for (type, _) in types.items():
    program.append(f"Visit{type}{base_name}(expr *{type}) (any, error)")
  program.append("}")


def define_ast(out: Path, base_name: str, types: dict[str, list[str]]):
  program = []
  
  program.append("package parser")
  program.append('import ( "mguzm4n/multichar-parser/src/lexer" )')
  
  program.append(f"type {base_name} interface {{")
  program.append("isExpr()")
  program.append("Accept (v Visitor) (any, error)")
  program.append("}")
  
  for (baseType, members) in types.items(): 
    
    # define struct
    program.append(f"type {baseType} struct {{")
    for member in members:
      name, memType = member.split(" ")
      program.append(f"{name} {memType}")
    program.append("}")
    
    # interface compliance
    program.append(f"func (*{baseType}) isExpr() {{ }}")
    
    # visitor pattern
    program.append(f"func ({baseType[0].lower()} *{baseType}) Accept(v Visitor) (any, error) {{")
    program.append(f"return v.Visit{baseType}{base_name}({baseType[0].lower()})")
    program.append("}")
    
    # constructor and fields
    cons_members = ", ".join(members)
    inst_members = ", ".join(
      list(map(lambda s: s.split(" ")[0], members))
    )
    program.append(f"func New{baseType}({cons_members}) *{baseType} {{")
    program.append(f"return &{baseType}{{")
    program.append(inst_members + ",")
    program.append("}")
    program.append("}")
    
  
  define_visitor(program, base_name, types)
  
  with open(out, 'w') as f:
    f.writelines(line + "\n" for line in program)
    
    
def main(out: Path):
  define_ast(
    out,
    "Expr", {
      "Binary": ["Left Expr", "Operator lexer.Token", "Right Expr"],
      "Grouping": ["Expression Expr"],
      "Literal": ["Value any"],
      "Unary": ["Operator lexer.Token", "Right Expr"]
    }
  )


if __name__ == "__main__":
  this_dir = Path(__file__).parent
  target_dir = this_dir.parent / "src" / "parser"
  if not target_dir.exists:
    raise Exception("directory not found")
    
  filename = "exprdefs.go"
  main(target_dir / filename)
