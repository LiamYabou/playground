class Person
    attr_accessor :name

    def initialize
        @name = "Liam"
    end
    
    def greeting
        "Hello, #{name}"
    end

    def profile
        "name: #{name}"
    end
end