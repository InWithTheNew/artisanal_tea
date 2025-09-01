// Utility to send form data to the backend
export interface SubmitPayload {
  Name: string;
  Command: string;
  User: string;
}

export async function submitFormData(payload: SubmitPayload): Promise<Response> {
  const url = process.env.REACT_APP_SUBMIT_URL;
  if (!url) throw new Error('REACT_APP_SUBMIT_URL is not set in environment variables');
  return fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(payload),
  });
}
