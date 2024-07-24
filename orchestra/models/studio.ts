type StudioResponse = {
  studio_id: number;
  subdomain: string;
  picture: string;
  name: string;
  description: string;
  address: string;
  city: string;
  zipcode: string;
  state: string;
  country: string;
  owner_id: number;
  created_at: string; // Use string for ISO date format
  updated_at: string; // Use string for ISO date format
  deleted_at?: string | null; // Use optional string or null for deleted_at
};
