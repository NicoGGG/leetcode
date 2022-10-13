export function longestPalindrome(s: string): string {
  let r = "";
  let l: number = s.length;
  for (let i = 0; i < l; i++) {
    for (let j = r.length + i + 1; j < l + 1; j++) {
      let tmp: string = s.slice(i, j);
      if (isPalindrome(tmp)) {
        r = tmp;
      }
    }
  }
  return r;
}

export function isPalindrome(s: string): boolean {
  const l = s.length;
  for (let i = 0; i < l; i++) {
    if (s[i] != s[l - i - 1]) {
      return false;
    }
  }
  return true;
}
