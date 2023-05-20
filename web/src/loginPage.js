import React, { useState } from "react";
import "./styles/loginPage.css";

function LoginPage() {
  // const [username, setUsername] = useState('');
  const [password, setPassword] = useState("");
  const [email, setMail] = useState("");

  return (
    <div className="bLogin">
      <a className="headText"> Войти в систему </a>
      <a className="email">E-mail: </a>
      <input
          type="email" value={email} name="en" onChange={(e) => setMail(e.target.value)}
      />
      <a className="password">Пароль: </a>
      <input
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
      />
      <button type="submit">Войти</button>
      <a className="newAcc">Нет акаунта?</a>
      <button type="submit">Зарегестрироваться</button>
    </div>
  );
}

export default LoginPage;
