export const isValidIsoDate = (str: string): boolean => {
  return /^\d{4,}-[01]\d-[0-3]\d/.test(str);
};

export const convertIsoDatesToDates = (obj: any) => {
  if (typeof obj === 'object' && obj !== null) {
    for (const key in obj) {
      if (obj.hasOwnProperty(key)) {
        const value = obj[key];
        if (typeof value === 'string' && isValidIsoDate(value)) {
          obj[key] = new Date(value);
        } else if (typeof value === 'object' && value !== null) {
          convertIsoDatesToDates(value); // Recursively check nested objects
        }
      }
    }
  }
  return obj;
};
