export async function fetchDropdownOptions(url: string): Promise<{ label: string; value: string }[]> {
  const res = await fetch(url);
  if (!res.ok) throw new Error('Failed to fetch dropdown options');
  const data = await res.json();
  // If the API returns a list of strings, convert to label/value pairs
  if (Array.isArray(data) && typeof data[0] === 'string') {
    return data.map((v: string) => ({ label: v, value: v }));
  }
  // Otherwise, assume it's already in the correct format
  return data;
}
