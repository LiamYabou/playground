Given(/^I am on the Google search page$/) do
    @driver = Selenium::WebDriver.for :chrome
    @driver.get "http://google.com"
end

When(/^I search for "([^"]*)"$/) do |keyword|
    puts "keyword: #{keyword}"
    element = @driver.find_element(name: "q")
    element.send_keys keyword
    element.submit
end

Then(/^the page title should start with "([^"]*)"$/) do |keyword|
    wait = Selenium::WebDriver::Wait.new(timeout: 10)
    wait.until { @driver.title.downcase.start_with? keyword }
    puts "Page title is #{@driver.title}"
    @driver.quit
end
