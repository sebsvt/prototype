import Cookies from "js-cookie";

const validateToken = async () => {
  const token = Cookies.get("access_token");
  if (!token) {
    return false;
  }

  try {
    const res = await fetch("http://localhost:8000/api/users/current_user", {
      method: "GET",
      headers: {
        Authorization: `${token}`,
      },
    });

    if (!res.ok) {
      throw new Error("Token is not valid");
    }

    console.log(res.body);
    return true;
  } catch (error) {
    console.error(error);
    return false;
  }
};

export default validateToken;
