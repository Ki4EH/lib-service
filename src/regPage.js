import React, { useState } from "react";

function AuthPage() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [usersurname, setUsersurname] = useState("");
  const [email, setemail] = useState("");
  const [checpassword, setchecPassword] = useState("");

  return (
    <center>
      <p class="probcenter" align="center">
        <div class="outline" align="center">
          <label>
            <h1>Регистрация</h1>

            <label>
              Имя:
              <br />
              <input
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                required
                placeholder="Ваше имя"
              />
            </label>
            <br />
            <label>
              Фамилия:
              <br />
              <input
                type="text"
                value={usersurname}
                onChange={(e) => setUsersurname(e.target.value)}
              />
            </label>
            <br />
            <label>
              E-mail:
              <br />
              <input
                type="email"
                value={email}
                name="en"
                onChange={(e) => setemail(e.target.value)}
              />
            </label>
            <br />
            <label>
              пароль:
              <br />
              <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </label>
            <br />
            <label>
              Повторите пароль:
              <br />
              <input
                type="password"
                value={checpassword}
                onChange={(e) => setchecPassword(e.target.value)}
              />
            </label>
            <br />
            <button type="submit">Зарегестрироваться</button>
          </label>
        </div>
        Есть акаунт?<button type="submit">Войти</button>
      </p>
    </center>
  );
}

export default AuthPage;
