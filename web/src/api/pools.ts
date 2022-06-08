const BASE_URL = import.meta.env.VITE_API_URL || "";

export const getPools = async () => {
  const response = await fetch(`${BASE_URL}/api/pools`);

  return await response.json();
};
