interface TokenResponse {
  email: string;
  exp: number;
  iss: string;
}

interface UserProfile {
  user_ref: number;
  firstname: string;
  surname: string;
  gender: string;
  phone: string;
  date_of_birth: string;
}
