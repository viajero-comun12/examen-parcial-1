from fastapi import FastAPI, HTTPException, status
from sqlmodel import SQLModel, Field 
from pydantic import EmailStr
from typing import Optional

class Usuario(SQLModel):
    id: Optional[int] = Field(default=None, primary_key=True)
    username: str = Field(index=True, unique=True)
    password: str
    email: Optional[EmailStr] = None
    is_active: bool = True

class UsuarioCrear(SQLModel):
    username: str
    password: str
    email: Optional[EmailStr] = None
    is_active: bool = True

class UsuarioLeer(SQLModel):
    id: int
    username: str
    email: Optional[EmailStr] = None
    is_active: bool

class UsuarioActualizar(SQLModel):
    username: Optional[str] = None
    email: Optional[EmailStr] = None
    is_active: Optional[bool] = None

class Logear(SQLModel):
    username: str
    password: str


usuario: list[Usuario] = [
    Usuario(id=1, username="admin", password="pacos", email="admin@example.com", is_active=True),
    Usuario(id=2, username="user", password="paco", email="user@example.com", is_active=True),
    Usuario(id=3, username="persona", password="pas", email="persona@example.com", is_active=False),
]
_next_id = 4  

app = FastAPI(title="API Login prebas")


def tenerid(user_id: int) -> Optional[Usuario]:
    for u in usuario:
        if u.id == user_id:
            return u
    return None

def teneruser(username: str) -> Optional[Usuario]:
    for u in usuario:
        if u.username == username:
            return u
    return None

def authenticate_user(username: str, password: str) -> bool:

    u = teneruser(username)
    if not u:
        return False
    if u.password != password:
        return False
    if not u.is_active:
        return False
    return True


@app.post("/users", response_model=UsuarioLeer)
def crear(payload: UsuarioCrear):
    if teneruser(payload.username) is not None:
        raise HTTPException(status_code=400, detail="El username ya existe")
    
    global _next_id
    nuevo = Usuario(
        id=_next_id,
        username=payload.username,
        password=payload.password,
        email=payload.email,
        is_active=payload.is_active,
    )
    usuario.append(nuevo)
    _next_id += 1
    return UsuarioLeer(id=nuevo.id, username=nuevo.username, email=nuevo.email, is_active=nuevo.is_active)


@app.get("/users", response_model=list[UsuarioLeer])
def listar():
    return [UsuarioLeer(id=u.id, username=u.username, email=u.email, is_active=u.is_active) for u in usuario]


@app.get("/user/id}", response_model=UsuarioLeer)
def obtener(user_id: int):
    u = tenerid(user_id)
    if not u:
        raise HTTPException(status_code=404, detail="Usuario no encontrado")
    return UsuarioLeer(id=u.id, username=u.username, email=u.email, is_active=u.is_active)


@app.put("/user/actualizar/id}", response_model=UsuarioLeer)
def actualizar(user_id: int, payload: UsuarioActualizar):
    u = tenerid(user_id)
    if not u:
        raise HTTPException(status_code=404, detail="Usuario no encontrado")
    
    if payload.username and payload.username != u.username:
        if teneruser(payload.username) is not None:
            raise HTTPException(status_code=400, detail="El usuario ya existe")
        u.username = payload.username
    if payload.email is not None:
        u.email = payload.email
    if payload.is_active is not None:
        u.is_active = payload.is_active

    return UsuarioLeer(id=u.id, username=u.username, email=u.email, is_active=u.is_active)


@app.delete("/user/borrar/id}")
def borrar(user_id: int):
    u = tenerid(user_id)
    if not u:
        raise HTTPException(status_code=404, detail="Usuario no encontrado")
    usuario.remove(u)
    return {"detail": "Usuario eliminado con éxito"}


@app.post("/login")
def logear(payload: Logear):
    ok = authenticate_user(payload.username, payload.password)
    if not ok:
        raise HTTPException(status_code=401, detail="Credenciales inválidas o usuario inactivo")
    return {"detail": "Acceso permitido. Bienvenido " + payload.username}

