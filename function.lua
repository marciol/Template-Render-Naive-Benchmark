f1 = function(table)
    return function(str)
        print(table)
        print(str)
    end
end

f1{ foo="bar" }[[baz, boo]]
