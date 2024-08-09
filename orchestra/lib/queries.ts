import axios from "axios";
import Cookies from "js-cookie";

const api = "http://localhost:8000/api";

const getUser = async (): Promise<UserProfile | null> => {
  try {
    const access_token = Cookies.get("access_token");
    if (!access_token) {
      throw new Error("Access token not found in cookies");
    }

    const verifyResponse = await axios.get<TokenResponse>(
      `${api}/auth/verify_token`,
      {
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      }
    );

    const userId = verifyResponse.data.iss;

    const profileResponse = await axios.get<UserProfile>(
      `${api}/profile/${userId}`,
      {
        headers: {
          Authorization: `Bearer ${access_token}`,
        },
      }
    );

    return profileResponse.data;
  } catch (error) {
    console.log("Error getting user profile:", error);
    return null;
  }
};

const getUserStudios = async (): Promise<StudioResponse[]> => {
  const access_token = Cookies.get("access_token");
  if (!access_token) {
    throw new Error("User not authenticated");
  }

  try {
    const response = await fetch("http://localhost:8000/api/studio/my-studio", {
      method: "GET",
      headers: {
        Authorization: `Bearer ${access_token}`,
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.error || "Failed to fetch studios");
    }

    const data: StudioResponse[] = await response.json();
    return data;
  } catch (error) {
    console.error("Error fetching studios:", error);
    throw error;
  }
};

const getStudioByDomain = async (domain: string): Promise<StudioResponse> => {
  try {
    const response = await fetch(`http://localhost:8000/api/studio/${domain}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.error || "Failed to fetch studio details");
    }

    const data: StudioResponse = await response.json();
    return data;
  } catch (error) {
    console.error("Error fetching studio details:", error);
    throw error;
  }
};

export { getUser, getUserStudios, getStudioByDomain };
