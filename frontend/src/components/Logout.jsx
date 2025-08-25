import { useEffect, useState } from "react";

const Logout = () => {
  const [msg, setMsg] = useState("");

  useEffect(() => {
    const logout = async () => {
      try {
        const res = await fetch("/logout", { 
            method: "POST",
            credentials: "include",
        });
        const data = await res.json();

        if (res.ok) {
          setMsg(data.message || "Logout successful!");
          localStorage.removeItem("lr_token");
          sessionStorage.removeItem("lr_token");

          setTimeout(() => {
            window.location.href = "/login";
          }, 1500);
        } else {
          setMsg(data.message || "Logout failed.");
        }
      } catch (err) {
        setMsg("Network error while logging out.");
      }
    };

    logout();
  }, []);

  return (
    <div style={{ textAlign: "center", marginTop: "50px" }}>
      <h2>{msg}</h2>
    </div>
  );
};

export default Logout;
