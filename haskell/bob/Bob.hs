module Bob (responseFor) where
  import Data.String.Utils
  import Data.Char

  responseFor :: String -> String
  responseFor request | strip request == "" = "Fine. Be that way!"
                      | (endswith "?" request) = "Sure."
                      | isShout request = "Woah, chill out!"
                      | otherwise = "Whatever."
                      where
                        isShout :: String -> Bool
                        isShout s = allUpper s && not (isNumeric s)

                        allUpper :: String -> Bool
                        allUpper s = s == map toUpper s

                        isNumeric :: String -> Bool
                        isNumeric s = all isNumber s
