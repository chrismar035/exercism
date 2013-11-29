module Bob (responseFor) where

responseFor :: String -> String
responseFor drivel | (last drivel) == "?" = "Sure."
                   | otherwise            = "Whatever."


