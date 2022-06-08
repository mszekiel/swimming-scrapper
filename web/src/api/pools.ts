export const getPools = async () => {
  const response = await fetch("/api/pools");

  return await response.json();
};
