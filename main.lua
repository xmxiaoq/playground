--
-- Created by IntelliJ IDEA.
-- User: 41956
-- Date: 2017/8/7
-- Time: 16:50
-- To change this template use File | Settings | File Templates.
--

json = require("dkjson")
require("collections")

function assert_tables_equal(tbl1, tbl2)
    if collect(tbl1):equals(tbl2) then
        return true
    end
    return error('Compared tables are not identical.')
end

function main()
--    print(json.encode({name="xiaoq", id=100}))
--    local obj2 = {'Taylor', 'php', 'javascript', 'lua'}
--    local obj = collect({name = 'Taylor', languages = {'php', 'javascript', 'lua'} }):flatten():all()
--    print(json.encode(obj))
--    print(assert_tables_equal(obj, obj2))
--
--    assert_tables_equal(
--        collect({name = 'Taylor', languages = {'php', 'javascript', 'lua'} }):flatten():all(),
--        {'Taylor', 'php', 'javascript', 'lua'}
--    )

    assert_tables_equal(
        collect({name = 'Liam', language = 'Lua'}):keys():all(),
        {'name', 'language'}
    )
end

main()