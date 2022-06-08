const BASE_URL = import.meta.env.VITE_API_URL || "";

interface Pool {
  identificator: number;
  location: { lat: number; lon: number };
  name: string;
}

interface CountedResponse<T> {
  count: number;
  data: T[];
}

export const getPools = async () => {
  const response = await fetch(`${BASE_URL}/api/pools`);

  return (await response.json()) as CountedResponse<Pool>;
};
