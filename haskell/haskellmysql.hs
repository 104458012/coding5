import Control.Monad
import Database.HDBC
import Database.HDBC.MySQL

main = do
  rows <- withRTSSignalsBlocked $ do
    conn <- connectMySQL defaultMySQLConnectInfo {
              mysqlHost = "localhost",
              mysqlUser = "mysql",
              mysqlPassword = "echoo710426"
              mysqlPort = "3360"
            }
quickQuery' conn "SELECT 1 + 1" []
