export function entryVisible(filter : string) {
  return (pokemon : string) => {
    return isEntryVisible(filter, pokemon);
  }
}

export function isEntryVisible(filter : string, pokemon : string) {
  if (!filter.trim()) {
    return true;
  }
  return pokemon.toLowerCase().indexOf(filter.toLowerCase()) !== -1;
}